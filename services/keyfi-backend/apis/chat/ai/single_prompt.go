package ai

import (
	"encoding/json"
	"fmt"
	"keyfi-backend/apis/util/interfaces"
	"keyfi-backend/models/chat/ai"
	"log"
	"net/http"
)

func SinglePrompt(w http.ResponseWriter, r *http.Request) {
	// Parse request
	var request ai.SinglePromptRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal("Failed to parse request: %v", err)
	}
	log.Printf("Incoming prompt: %s\n", request.Prompt)

	// Format response
	response := interfaces.SendTextPrompt(request.Prompt)
	partString := ""
	for i := range response.Candidates[0].Content.Parts {
		partString = fmt.Sprintf("%s%s", partString, response.Candidates[0].Content.Parts[i])
	}
	fmt.Printf("response: %s\n", partString)

	// Return result
	reply := ai.SinglePromptResponse{
		Response: partString,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(reply)
	if err != nil {
		log.Fatal("Failed to encode response: %v", err)
	}
}