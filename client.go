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
	send chan Message
}

func (c *Client) Read() {
	defer c.conn.Close()
	for {
		_, m, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		var msg Message
		if err = msg.Unmarshal(m); err != nil {
			return
		}

		c.channel.Consume(msg)
	}
}

func (c *Client) Write() {
	defer c.conn.Close()
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg.Bytes())
		if err != nil {
			return
		}
	}
}
