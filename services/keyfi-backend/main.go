package main

import (
	"keyfi-backend/apis/chat/ai"
	pb "keyfi-backend/protos"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

const GOOGLE_AI_API_PATH = "./googleai_apikey"

func main() {
	// Read API key for Google AI
	dat, err := os.ReadFile(GOOGLE_AI_API_PATH)
	if err != nil {
		panic(err)
	}
	os.Setenv("GOOGLEAI_API_KEY", string(dat))

	// set a port for the server
	port := ":8080"

	// listen for requests on 8080
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKeyFiAIServiceServer(grpcServer, &ai.Server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}

}