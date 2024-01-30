package main

// Channel represents a text channel for chat.
type Channel struct {
	// inbound holds incoming messages.
	inbound chan Message

	// enter is for clients entering the room.
	enter chan *Client

	// exit is for clients exiting the room.
	exit chan *Client

	// clients holds registered clients.
	clients map[*Client]bool
}

// NewChannel makes a new channel
func NewChannel() *Channel {
	return &Channel{
		inbound: make(chan Message),
		enter:   make(chan *Client),
		exit:    make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

func (c *Channel) Run() {
	for {
		select {
		case client := <-c.enter:
			c.clients[client] = true
			for client := range c.clients {
				client.counter <- c.Size()
			}
		case client := <-c.exit:
			delete(c.clients, client)
			close(client.send)
			for client := range c.clients {
				client.counter <- c.Size()
			}
		case msg := <-c.inbound:
			for client := range c.clients {
				client.send <- msg
			}
		}
	}
}

func (c *Channel) Enter(client *Client) {
	c.enter <- client
}

func (c *Channel) Exit(client *Client) {
	c.exit <- client
}

func (c *Channel) Consume(message Message) {
	c.inbound <- message
}

func (c *Channel) Size() UserCounter {
	return UserCounter(len(c.clients))
}
