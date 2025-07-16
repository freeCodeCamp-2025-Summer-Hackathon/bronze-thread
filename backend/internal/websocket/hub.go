package websocket

import "github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/models"

// Message represents a message to be broadcasted to a chat room.
type Message struct {
	RoomID  uint
	UserID  uint
	Content string
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients in the correct rooms.
type Hub struct {
	// Registered clients. Maps client pointer to boolean.
	clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan models.ChatMessage

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	// A map of chat rooms. The key is the room ID.
	rooms map[uint]map[*Client]bool
}

// NewHub creates a new Hub instance.
func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan models.ChatMessage),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		rooms:      make(map[uint]map[*Client]bool),
	}
}

// Run starts the hub's event loop.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = true
			if h.rooms[client.RoomID] == nil {
				h.rooms[client.RoomID] = make(map[*Client]bool)
			}
			h.rooms[client.RoomID][client] = true

		case client := <-h.Unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				if room, ok := h.rooms[client.RoomID]; ok {
					delete(room, client)
					if len(room) == 0 {
						delete(h.rooms, client.RoomID)
					}
				}
				close(client.Send)
			}

		case message := <-h.Broadcast:
			if room, ok := h.rooms[message.ChatRoomID]; ok {
				for client := range room {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(h.clients, client)
						delete(room, client)
						if len(room) == 0 {
							delete(h.rooms, client.RoomID)
						}
					}
				}
			}
		}
	}
}
