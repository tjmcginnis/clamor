package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/websocket"
	"github.com/tjmcginnis/namer"
)

var (
	addr         = flag.String("addr", ":8080", "The address of the server.")
	homeTemplate = filepath.Join("templates", "index.html")
	upgrader     = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	flag.Parse()

	c := NewChannel()

	homeTempl := template.Must(template.ParseFiles(homeTemplate))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homeTempl.Execute(w, r)
	})

	http.Handle("/channel", &chatHandler{channel: c})

	go c.Run()

	log.Println("Starting server on: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type chatHandler struct {
	channel *Channel
}

func (c *chatHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP: ", err)
		return
	}

	namer := namer.New()
	client := &Client{
		channel: c.channel,
		conn:    conn,
		send:    make(chan Message),
		user:    NewUser(namer.Name()),
	}
	c.channel.Enter(client)
	defer func() { c.channel.Exit(client) }()
	go client.Write()
	client.Read()
}
