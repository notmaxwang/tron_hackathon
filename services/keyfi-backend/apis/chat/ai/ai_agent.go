package ai

import (
	"context"
	"fmt"
	"keyfi-backend/apis/util/interfaces"
	pb "keyfi-backend/protos/ai"
	"log"
)

type Server struct {
	pb.UnimplementedAIServiceServer
}

func (s *Server) SinglePrompt(ctx context.Context, request *pb.SinglePromptRequest) (*pb.SinglePromptResponse, error) {
	log.Printf("Incoming prompt: %s\n", request.GetPrompt())

	response := interfaces.SendTextPrompt(request.GetPrompt())
	partString := ""
	for i := range response.Candidates[0].Content.Parts {
		partString = fmt.Sprintf("%s%s", partString, response.Candidates[0].Content.Parts[i])
	}
	log.Printf("response: %s\n", partString)

	return &pb.SinglePromptResponse{Response: partString}, nil
}
