package main

import (
	"keyfi-backend/apis/chat/ai"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const GOOGLE_AI_API_PATH = "./googleai_apikey"

func main() {
	// Read API key for Google AI
	dat, err := os.ReadFile(GOOGLE_AI_API_PATH)
	if err != nil {
		panic(err)
	}
	os.Setenv("GOOGLEAI_API_KEY", string(dat))

	router := mux.NewRouter()

	router.HandleFunc("/singlePrompt", ai.SinglePrompt).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
