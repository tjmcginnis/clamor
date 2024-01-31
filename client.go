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

	// counter tracks the number of users in the current channel.
	counter chan UserCounter

	user *User
}

func NewClient(channel *Channel, conn *websocket.Conn, user *User) *Client {
	return &Client{
		channel: channel,
		conn:    conn,
		send:    make(chan Message),
		counter: make(chan UserCounter),
		user:    user,
	}
}

func (c *Client) SendMessage() {
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
		msg.User = c.user

		c.channel.Consume(msg)
	}
}

func (c *Client) ReceiveMessages() {
	defer c.conn.Close()
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg.Bytes())
		if err != nil {
			return
		}
	}
}

func (c *Client) UpdateCounter() {
	defer c.conn.Close()
	for counter := range c.counter {
		err := c.conn.WriteMessage(websocket.TextMessage, counter.Bytes())
		if err != nil {
			return
		}
	}
}
