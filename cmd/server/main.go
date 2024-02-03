package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/websocket"
	"github.com/tjmcginnis/namer"

	"github.com/tjmcginnis/clamor"
)

var (
	addr          = flag.String("addr", ":8080", "The address of the server.")
	templateFiles = []string{
		filepath.Join("templates", "index.html"),
		filepath.Join("templates", "user_counter.html"),
	}
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	flag.Parse()

	c := clamor.NewChannel()

	templates, err := template.ParseFiles(templateFiles...)
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Host    string
			Counter clamor.UserCounter
		}{
			Host:    r.Host,
			Counter: c.Size().Increment(),
		}
		templates.ExecuteTemplate(w, "index", &data)
	})

	http.Handle("/channel", &chatHandler{
		channel: c,
		namer:   namer.New(),
	})

	go c.Run()

	log.Println("Starting server on: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type chatHandler struct {
	channel *clamor.Channel
	namer   namer.Namer
}

func (c *chatHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP: ", err)
		return
	}

	name := c.namer.Name()
	client := clamor.NewClient(c.channel, conn, clamor.NewUser(name))

	c.channel.Enter(client)
	defer func() { c.channel.Exit(client) }()
	go client.ReceiveMessages()
	go client.UpdateCounter()

	client.SendMessage()
}
