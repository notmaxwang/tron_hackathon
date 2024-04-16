package chat

import (
	"keyfi-backend/util/persistence/models"
	"keyfi-backend/util/persistence"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	Upgrader websocket.Upgrader
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Configure WebSocket upgrader
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Implement your origin check logic here
			// For example, allow connections only from example.com
			return true // Should change this to our own VPC or something.
		},
	}

	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer conn.Close()

	// WebSocket connection handling
	for {
		// Read message from WebSocket client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read:", err)
			break
		}
		log.Printf("Received message: %s\n", msg)
		// Echo message back to client
		suffix := " idk"
		err = conn.WriteMessage(websocket.TextMessage, append(msg, suffix...))
		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}
