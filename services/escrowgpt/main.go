package main

import (
	"context"
	"escrowgpt/interfaces"
	"protos"
	pb "escrowgpt/protos"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

const GOOGLE_AI_API_PATH = "./.googleai_apikey"

// an RPC server in Go

type server struct {
	protos.UnimplementedEscrowGPTServiceServer
}

func (s *server) SinglePrompt(ctx context.Context, request *pb.SinglePromptRequest) (*pb.SinglePromptResponse, error) {
	log.Printf("Incoming prompt: %s\n", request.GetPrompt())

	response := interfaces.SendTextPrompt(request.GetPrompt())
	partString := ""
	for i := range response.Candidates[0].Content.Parts {
		partString = fmt.Sprintf("%s%s", partString, response.Candidates[0].Content.Parts[i])
	}
	fmt.Printf("response: %s\n", partString)

	return &pb.SinglePromptResponse{Response: partString}, nil
}

// func (t *BannerMessageServer) GetBannerMessage(args *Args, reply *string) error {
// 	fmt.Printf("prompt: %s\n", args.Message)

// 	response := interfaces.SendTextPrompt(args.Message)
// 	partString := ""
// 	for i := range response.Candidates[0].Content.Parts {
// 		partString = fmt.Sprintf("%s%s", partString, response.Candidates[0].Content.Parts[i])
// 	}
// 	fmt.Printf("response: %s\n", partString)
// 	*reply = partString
// 	return nil
// }

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

	s := grpc.NewServer()
	pb.RegisterEscrowGPTServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}

}
