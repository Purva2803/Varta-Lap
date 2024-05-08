package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ChatRoom struct {
	clients map[*websocket.Conn]bool
}

func (cr *ChatRoom) broadcast(messageType int, message []byte) {
	for client := range cr.clients {
		err := client.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error broadcasting message to client:", err)
			client.Close()
			delete(cr.clients, client)
		}
	}
}

func (cr *ChatRoom) join(client *websocket.Conn) {
	cr.clients[client] = true
}

func (cr *ChatRoom) leave(client *websocket.Conn) {
	if _, ok := cr.clients[client]; ok {
		delete(cr.clients, client)
		client.Close()
	}
}

func main() {
	chatRooms := make(map[string]*ChatRoom)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		room := r.URL.Query().Get("room")
		if room == "" {
			http.Error(w, "Room parameter is required", http.StatusBadRequest)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error upgrading to WebSocket:", err)
			return
		}

		if _, ok := chatRooms[room]; !ok {
			chatRooms[room] = &ChatRoom{clients: make(map[*websocket.Conn]bool)}
		}

		chatRoom := chatRooms[room]
		chatRoom.join(conn)

		defer chatRoom.leave(conn)

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message from WebSocket:", err)
				break
			}
			chatRoom.broadcast(messageType, message)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
