package chat

import (
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

	// Start connection to AI
	convo, starter, err := StartConvo()
	if err != nil {
		log.Println("Error while trying to establish AI convo session", err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte(starter))
	if err != nil {
		log.Println("Write:", err)
		return
	}

	// WebSocket connection handling
	for {
		// Read message from WebSocket client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read:", err)
			break
		}
		log.Printf("Received message: %s\n", msg)

		reply, err := convo.SendChatPrompt(string(msg))
		if err != nil {
			log.Println("error while trying to send prompt", err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(reply))
		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}
