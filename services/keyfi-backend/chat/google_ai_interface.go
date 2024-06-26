package chat

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Conversation struct {
	session *genai.ChatSession
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

func StartConvo(ctx context.Context) (*Conversation, string, error) {
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_AI_KEY")))
	if err != nil {
		return nil, "", err
	}
	defer client.Close()

	// For text-and-image input (multimodal), use the gemini-pro-vision model
	model := client.GenerativeModel("gemini-pro")
	cs := model.StartChat()

	// promptContext, err = ioutil.ReadFile("./prompt_context.txt")
	// if err != nil {
	// 	log.Println("error while reading prompt context file", err)
	// 	return nil, err
	// }
	promptContext := "you are a real estate chatbot, you will answer the user's questions about real estate. now, please greet the user with a warm welcoming"

	prompt := genai.Text(promptContext)
	resp, err := cs.SendMessage(ctx, prompt)
	if err != nil {
		log.Println("couldn't send first prompt to gemini", err)
		return nil, "", err
	}

	fullResponse := ""
	for i := range resp.Candidates[0].Content.Parts {
		fullResponse = fmt.Sprintf("%s%s", fullResponse, resp.Candidates[0].Content.Parts[i])
	}

	return &Conversation{
		session: cs,
	}, fullResponse, nil
}

func (convo *Conversation) SendChatPrompt(ctx context.Context, promptString string) (string, error) {
	prompt := genai.Text(promptString)
	resp, err := convo.session.SendMessage(ctx, prompt)
	if err != nil {
		log.Println("couldn't send first prompt to gemini", err)
		return "", err
	}

	fullResponse := ""
	for i := range resp.Candidates[0].Content.Parts {
		fullResponse = fmt.Sprintf("%s%s", fullResponse, resp.Candidates[0].Content.Parts[i])
	}

	return fullResponse, nil
}