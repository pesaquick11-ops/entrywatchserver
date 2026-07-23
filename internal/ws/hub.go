package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	mu      sync.RWMutex
	clients map[*websocket.Conn]bool
}

func NewHub() *Hub {
	return &Hub{clients: make(map[*websocket.Conn]bool)}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: lock this down to your Next.js origin in production
		return true
	},
}

// ServeWS is an http.HandlerFunc you register on your router
func (h *Hub) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ws upgrade error:", err)
		return
	}
	h.addClient(conn)
	log.Println("🔌 client connected:", conn.RemoteAddr())

	go func() {
		defer h.removeClient(conn)
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				log.Println("🔌 client disconnected:", conn.RemoteAddr())
				return
			}
		}
	}()
}

func (h *Hub) addClient(c *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[c] = true
}

func (h *Hub) removeClient(c *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, c)
	c.Close()
}

// Broadcast sends any JSON-serializable payload to all connected clients
func (h *Hub) Broadcast(v interface{}) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for c := range h.clients {
		if err := c.WriteJSON(v); err != nil {
			log.Println("write error:", err)
			c.Close()
			delete(h.clients, c)
		}
	}
}
