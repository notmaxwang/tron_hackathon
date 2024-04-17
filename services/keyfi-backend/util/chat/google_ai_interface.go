package chat

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Conversation struct {
	chatIter genai.GenerateContentResponseIterator
}

var once sync.Once

func SendTextPrompt(message string) *genai.GenerateContentResponse {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_AI_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	resp, err := model.GenerateContent(ctx, genai.Text(message))
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func StartConvo() (*Conversation, error) {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	log.Printf("key: %s\n", os.Getenv("GEMINI_AI_KEY"))
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_AI_KEY")))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// For text-and-image input (multimodal), use the gemini-pro-vision model
	model := client.GenerativeModel("gemini-pro")

	prompt := genai.Text("Tell me a story about this animal")
	return &Conversation{
		chatIter: *model.GenerateContentStream(ctx, prompt),
	}, nil
}

func (convo *Conversation) SendChatPrompt(prompt string) (*genai.GenerateContentResponse, error) {
	resp, err := convo.chatIter.Next()
	if err == iterator.Done {
		log.Printf("the chat has ended\n")
		return nil, nil
	}

	return resp, err
}
