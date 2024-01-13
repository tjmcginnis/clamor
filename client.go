package main

import (
	"github.com/gorilla/websocket"
)

// Client represents a chat user.
type Client struct {
	channel *Channel

	// conn is the websocket connection.
	conn *websocket.Conn

	// send is a channel for messages.
	send chan []byte
}

func (c *Client) Read() {
	defer c.conn.Close()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		c.channel.Consume(msg)
	}
}

func (c *Client) Write() {
	defer c.conn.Close()
	for sent := range c.send {
		var msg Message
		err := msg.Unmarshal(sent)
		if err != nil {
			return
		}

		err = c.conn.WriteMessage(websocket.TextMessage, msg.Bytes())
		if err != nil {
			return
		}
	}
}
