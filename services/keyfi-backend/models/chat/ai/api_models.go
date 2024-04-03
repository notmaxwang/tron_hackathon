package ai

type SinglePromptRequest struct {
	Prompt string `json:"prompt"`
}

type SinglePromptResponse struct {
	Response string `json:"response"`
}
