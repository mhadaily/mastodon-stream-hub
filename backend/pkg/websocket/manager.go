package websocketmanager

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan []byte)
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Printf("WebSocket error: %v", err)
			delete(clients, ws)
			break
		}
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.BinaryMessage, msg)
			if err != nil {
				log.Printf("Error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func BroadcastMessage(message []byte) {
	broadcast <- message
}
