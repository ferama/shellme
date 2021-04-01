package wshub

import (
	"shellme/logger"

	"github.com/gorilla/websocket"
)

const (
	clientLogContext = "[WS-CLIENT]"
)

// Client ...
type Client struct {
	conn *websocket.Conn
	send chan HubMessage
	quit chan bool

	OnDisconnect func()
}

// NewClient ...
func NewClient(conn *websocket.Conn) *Client {
	client := &Client{
		conn: conn,
		quit: make(chan bool),
		send: make(chan HubMessage),
	}

	return client
}

func (c *Client) close() {
	c.conn.Close()
}

// Serve ...
func (c *Client) Serve() {
	go c.watchDisconnect()

	for {
		select {
		// case <-c.quit:
		// 	if err := c.conn.Close(); err != nil {
		// 		logger.ErrorWithFields(clientLogContext, "ws client connection error", logger.Fields{
		// 			"err": err,
		// 		})
		// 	}
		// 	logger.Info(clientLogContext, "WS Client quit")

		// 	return
		case msg := <-c.send:
			if err := c.conn.WriteJSON(msg); err != nil {
				return
			}
		}
	}
}

func (c *Client) watchDisconnect() {
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			c.OnDisconnect()
			// if websocket.IsCloseError(err) {
			// 	logger.Info(clientLogContext, "client disconnected")
			// }
			logger.Info(clientLogContext, "Client disconnected")
			close(c.send)
			// close(c.quit)
			return
		}
	}
}
