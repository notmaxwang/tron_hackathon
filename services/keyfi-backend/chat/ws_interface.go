package chat

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/gorilla/websocket"
	"google.golang.org/api/option"
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
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_AI_KEY")))
	if err != nil {
		log.Println("failed to create new client", err)
		return
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.0-pro")
	cs := model.StartChat()

	sendMessageToGemini := func(msg string) (string, error) {
		fmt.Printf("== Me: %s\n== Model:\n", msg)
		res, err := cs.SendMessage(ctx, genai.Text(msg))
		if err != nil {
			log.Println("failed to send message", err)
			return "", err
		}

		result := ""
		for i := range res.Candidates[0].Content.Parts {
			result = fmt.Sprintf("%s%s", result, res.Candidates[0].Content.Parts[i])
		}
		return result, nil
	}

	// Read file into a string
	starterPrompt, err := ioutil.ReadFile("./chat/prompt_context.txt")
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}

	starter, err := sendMessageToGemini(string(starterPrompt))
	if err != nil {
		log.Println("send message failure", err)
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

		reply, err := sendMessageToGemini(string(msg))
		if err != nil {
			log.Println("send message failure", err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(reply))
		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}
