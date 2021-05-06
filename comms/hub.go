package comms

var GlobalHub *Hub

type ClientBroadcast struct {
	Client	*Client
	Stream	[]byte
}
// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan ClientBroadcast

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan ClientBroadcast),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) SendToAll(bs []byte) {
	for client := range h.clients {
		select {
		case client.send <- bs:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}
}

func (h *Hub) SendTo(client *Client, bs []byte) {
	select {
	case client.send <- bs:
	default:
		close(client.send)
		delete(h.clients, client)
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			h.SendToAll(message.Stream)
		}
	}
}
