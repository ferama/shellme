package wshub

import (
	"shellme/logger"
	"sync"
)

const (
	logContext = "[WSHUB]"
)

// HubSession ...
type HubSession struct {
	Name   string
	Client *Client
	Wg     *sync.WaitGroup
}

// HubMessage ...
type HubMessage struct {
	SessionName string `json:"sessionName"`
	CellID      int    `json:"cellID"`
	RunID       int    `json:"runID"`
	Data        string `json:"data"`
	Channel     string `json:"channel"`
	ContentType string `json:"contentType"`
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]map[*Client]bool

	// Inbound messages for the clients.
	Broadcast chan HubMessage

	// Register clients.
	Register chan *HubSession
	// Unregister clients.
	Unregister chan *HubSession
	// destroy session by name
	DestroySession chan string
}

var hubInstance *Hub

// GetHubInstance ...
func GetHubInstance() *Hub {
	if hubInstance != nil {
		return hubInstance
	}

	hubInstance = &Hub{
		Broadcast:      make(chan HubMessage),
		Register:       make(chan *HubSession),
		Unregister:     make(chan *HubSession),
		DestroySession: make(chan string),
		clients:        make(map[string]map[*Client]bool),
	}
	return hubInstance
}

// Run ...
func (h *Hub) Run() {
	logger.Info(logContext, "Starting ws hub")
	for {
		select {
		case name := <-h.DestroySession:
			logger.Info(logContext, "Destroying session "+name)

			for client := range h.clients[name] {
				delete(h.clients[name], client)
				// close(client.send)
				client.close()
			}
			delete(h.clients, name)

		case s := <-h.Register:
			if _, ok := h.clients[s.Name]; !ok {
				h.clients[s.Name] = make(map[*Client]bool)
			}
			h.clients[s.Name][s.Client] = true
			logger.Info(logContext, "Registered client for session "+s.Name)

		case s := <-h.Unregister:
			if _, ok := h.clients[s.Name][s.Client]; ok {
				delete(h.clients[s.Name], s.Client)
				// s.Client.Close()
				logger.Info(logContext, "Unregistered client for session "+s.Name)
			}
			s.Wg.Done()

		case message := <-h.Broadcast:
			// log.Printf("sending %v %v: %v", len(h.clients[message.SessionName]), message.SessionName, message.Data)
			for client := range h.clients[message.SessionName] {
				select {
				case client.send <- message:
					// default:
					// client.Close()
					// delete(h.clients[message.SessionName], client)
				}
			}
		}
	}
}
