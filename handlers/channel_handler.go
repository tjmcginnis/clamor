package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tjmcginnis/namer"

	"github.com/tjmcginnis/clamor"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ChannelHandler struct {
	Channel *clamor.Channel
	Namer   namer.Namer
}

func (h ChannelHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP: ", err)
		return
	}

	name := h.Namer.Name()
	client := clamor.NewClient(h.Channel, conn, clamor.NewUser(name))

	h.Channel.Enter(client)
	defer func() { h.Channel.Exit(client) }()
	go client.ReceiveMessages()
	go client.UpdateCounter()

	client.SendMessage()
}
