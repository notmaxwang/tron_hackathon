package ai

import (
	"context"
	"os"
	pb "keyfi-backend/protos/query"
)

type Server struct {
	pb.UnimplementedQueryServiceServer
}

func (s *Server) GetValues(ctx context.Context, request *pb.GetValuesRequest) (*pb.GetValuesResponse, error) {

	keys := request.Keys
	numKeys := len(keys)
	result := make([]*pb.KeyValuePair, numKeys)

	for i:=0;i<numKeys;i++ {
		value := os.Getenv(keys[i])

		result[i] = &pb.KeyValuePair{
			Key: keys[i],
			Value: value,
		}
	}

	return &pb.GetValuesResponse{KeyValuePairs: result}, nil
}
