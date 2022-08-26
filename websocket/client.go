package websocket

import "github.com/gorilla/websocket"

type Client struct {
	Hub      *Hub
	Id       string
	Socket   *websocket.Conn
	OutBound chan []byte
}

func NewClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		Hub:      hub,
		Socket:   socket,
		OutBound: make(chan []byte),
	}
}

func (c *Client) Write() {
	for {
		select {
		case msg, ok := <-c.OutBound:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
func (c Client) Close() {
	c.Socket.Close()
	close(c.OutBound)
}