package ai

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

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
