package main

import (
	"context"
	"log"

	pb "goclient/protos" // Import generated protobuf code

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Create a gRPC client
	client := pb.NewKeyFiAIServiceClient(conn)

	// Contact the server and print out its response.
	response, err := client.SinglePrompt(context.Background(), &pb.SinglePromptRequest{Prompt: "yaobin is a nice guy :)"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", response.Response)
}
