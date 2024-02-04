package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/tjmcginnis/namer"

	"github.com/tjmcginnis/clamor"
	"github.com/tjmcginnis/clamor/handlers"
)

var addr = flag.String("addr", ":8080", "The address of the server.")

func main() {
	flag.Parse()

	channel := clamor.NewChannel()
	counter := clamor.NewUserCounter()

	http.Handle("/", handlers.ChatHandler{Counter: &counter})
	http.Handle("/channel", handlers.ChannelHandler{
		Channel: channel,
		Namer:   namer.New(),
	})

	go channel.Run()

	log.Println("Starting server on: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
