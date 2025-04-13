package controllers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Permite conexiones de cualquier origen (ajusta para producción)
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
type Client struct {
	conn *websocket.Conn
	send chan []byte
}
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}