package main

import (
	"context"
	"log"

	pb "goclient/protos/ai" // Import generated protobuf code
	pb2 "goclient/protos/query"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("34.236.81.43:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Create a gRPC client
	client := pb.NewAIServiceClient(conn)

	// Contact the server and print out its response.
	response, err := client.SinglePrompt(context.Background(), &pb.SinglePromptRequest{Prompt: "yaobin is a nice guy :)"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", response.Response)

	client2 := pb2.NewQueryServiceClient(conn)

	// Contact the server and print out its response.
	keys := []string{"GOOGLE_MAPS_KEY", "Yaobin"}
	response2, err2 := client2.GetValues(context.Background(), &pb2.GetValuesRequest{Keys: keys})
	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}
	log.Printf("Response: %v", response2.KeyValuePairs)
}
