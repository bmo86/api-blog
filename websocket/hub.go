package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	Clients    []*Client
	Register   chan *Client
	Unregister chan *Client
	Mutex      *sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make([]*Client, 0),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Mutex:      &sync.Mutex{},
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case c := <-hub.Register:
			hub.onConnect(c)
		case c := <-hub.Unregister:
			hub.onDisconnects(c)
		}
	}
}

func (hub *Hub) Broadcast(msg interface{}, ingnore *Client) {

	data, _ := json.Marshal(msg)
	for _, client := range hub.Clients {
		if client != ingnore {
			client.OutBound <- data
		}

	}

}

func (hub *Hub) onConnect(client *Client) {
	log.Println("Client connected", client.Socket.RemoteAddr())

	hub.Mutex.Lock()
	defer hub.Mutex.Unlock()
	client.Id = client.Socket.RemoteAddr().String()
	hub.Clients = append(hub.Clients, client)
}

func (hub *Hub) onDisconnects(client *Client) {
	log.Println("Client Disconnect", client.Socket.RemoteAddr())
	client.Close()
	hub.Mutex.Lock()
	defer hub.Mutex.Unlock()
	i := -1
	for j, c := range hub.Clients {
		if c.Id == client.Id {
			i = j
			break
		}
	}

	copy(hub.Clients[i:], hub.Clients[i+1:])
	hub.Clients[len(hub.Clients)-1] = nil
	hub.Clients = hub.Clients[:len(hub.Clients)-1]
}



func (hub *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error upgrading connection", http.StatusInternalServerError)
	}

	client := NewClient(hub, socket)
	hub.Register <- client
	
	go client.Write()
}