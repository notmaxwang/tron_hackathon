package ai

import (
	"encoding/json"
	"fmt"
	"keyfi-backend/apis/util/interfaces"
	"keyfi-backend/models/chat/ai"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SinglePrompt(w http.ResponseWriter, r *http.Request) {
	// Parse request
	params := mux.Vars(r)
	prompt := params["prompt"]

	// Format response
	response := interfaces.SendTextPrompt(prompt)
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
	err := json.NewEncoder(w).Encode(reply)
	if err != nil {
		log.Fatal("Failed to encode response: %v", err)
	}
}
