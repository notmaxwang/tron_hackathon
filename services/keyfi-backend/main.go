package main

import (
	"keyfi-backend/apis/chat/ai"
	"log"
	"net/http"
	"os"
)

const GOOGLE_AI_API_PATH = "./googleai_apikey"

func main() {
	// Read API key for Google AI
	dat, err := os.ReadFile(GOOGLE_AI_API_PATH)
	if err != nil {
		panic(err)
	}
	os.Setenv("GOOGLEAI_API_KEY", string(dat))

	http.HandleFunc("/simplePrompt", ai.SinglePrompt)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
