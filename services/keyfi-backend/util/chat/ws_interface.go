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
	convo, err := ai.StartConvo()
	if err != nil {
		log.Println("Error while trying to establish AI convo session")
		return nil, err
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

		aiResponse, err := convo.SendChatPrompt(msg)
		if aiResponse == nil && err == nil {
			log.Println("chat ended for some reason")
		}
		if err != nil {
			log.Println("error while trying to send prompt")
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, aiResponse)
		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}
