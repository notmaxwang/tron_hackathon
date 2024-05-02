package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"keyfi-backend/util/persistence"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/google/generative-ai-go/genai"
	"github.com/gorilla/websocket"
	"google.golang.org/api/option"
)

var (
	userMessages = make(chan Message)
	geminiMessages = make(chan Message)
	backendMessages = make(chan Message)
)

type WebSocketHandler struct {
	Upgrader websocket.Upgrader
}

type Message struct {
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	Message  string    `json:"message"`
	Commands []Command `json:"commands"`
	Results  []Result  `json:"results"`
}

type Command struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type Result struct {
	Command string   `json:"command"`
	Results []string `json:"results"`
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

	starterJson, err := sendMessageToGemini(string(starterPrompt))
	if err != nil {
		log.Println("send message failure", err)
		return
	}

	// Create a variable of type Message to store the parsed JSON
	var geminiMsg Message
	// Parse the JSON data into the 'msg' variable
	err = json.Unmarshal([]byte(starterJson), &geminiMsg)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return
	}

	if geminiMsg.Receiver != "user" {
		log.Println("probably some error with prompt because the first message not sent to user")
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte(geminiMsg.Message))
	if err != nil {
		log.Println("Write:", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(3)

	// Create separate goroutines to handle messages for each listener
	go func() {
		handleBackend(conn)
	}()
	go func() {
		handleGemini(conn, cs, ctx)
	}()
	go func() {
		handleUser(conn)
	}()

	// WebSocket connection handling
	for {
		// Read message from WebSocket client
		_, msgRaw, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read:", err)
			break
		}

		msg := Message{
			Sender:   "user",
			Receiver: "gemini",
			Message:  string(msgRaw),
			Commands: []Command{},
			Results:  []Result{},
		}

		sendMessage(msg)
	}
}

func handleGemini(conn *websocket.Conn, cs *genai.ChatSession, ctx context.Context) {
	for msg := range geminiMessages {
		log.Println("gemini saw new message")
		if msg.Receiver == "gemini" {
			// Convert the Message struct to JSON
			msgJson, err := json.Marshal(msg)
			if err != nil {
				log.Println("Error marshalling JSON:", err)
				return
			}

			res, err := cs.SendMessage(ctx, genai.Text(msgJson))
			if err != nil {
				log.Println("failed to send message", err)
				return
			}

			reply := ""
			for i := range res.Candidates[0].Content.Parts {
				reply = fmt.Sprintf("%s%s", reply, res.Candidates[0].Content.Parts[i])
			}

			var geminiMsg Message
			err = json.Unmarshal([]byte(reply), &geminiMsg)
			if err != nil {
				log.Println("Error parsing JSON:", err)
				return
			}

			sendMessage(geminiMsg)
		}
	}
	log.Println("gemini listener closed")
}

func handleBackend(conn *websocket.Conn) {
	dao, err := persistence.GetListingsDao()
	if err != nil {
		log.Println("could not access DB", err)
		return
	}

	for msg := range backendMessages {
		log.Println("backend saw new message")
		if msg.Receiver == "backend" {
			responseMsg := Message{
				Sender:   "backend",
				Receiver: "gemini",
				Message:  "",
				Commands: []Command{},
				Results:  []Result{},
			}
			for _, command := range msg.Commands {
				if command.Command == "getListingDetailsFromSqlQuery" {
					if len(command.Args) != 1 {
						responseMsg.Results = append(responseMsg.Results, Result{
							Command: command.Command,
							Results: []string{"ERROR: expected 1 single SQL query as string"},
						})
					}

					results, err := dao.QueryForAI(command.Args[0])
					if err != nil {
						responseMsg.Results = append(responseMsg.Results, Result{
							Command: command.Command,
							Results: []string{"error querying DB"},
						})
					}

					responseMsg.Results = append(responseMsg.Results, Result{
						Command: command.Command,
						Results: *results,
					})
				} else {
					responseMsg.Results = append(responseMsg.Results, Result{
						Command: command.Command,
						Results: []string{"UNSUPPORTED_COMMAND"},
					})
				}
			}
			sendMessage(responseMsg)
		}
	}
}

func handleUser(conn *websocket.Conn) {
	for msg := range userMessages {
		if msg.Receiver == "user" {
			// Send message to client
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Message))
			if err != nil {
				log.Println("Error sending message:", err)
				return
			}
		}
	}
}

func sendMessage(msg Message) {
	log.Println(msg)
	switch msg.Receiver {
	case "user":
		userMessages <- msg
	case "backend":
		backendMessages <- msg
	case "gemini":
		geminiMessages <- msg
	}
}
