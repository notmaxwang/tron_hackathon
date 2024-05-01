package main

import (
	"context"
	"log"

	pb "goclient/protos/listing" // Import generated protobuf code
	pb2 "goclient/protos/query"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("ec2-34-236-81-43.compute-1.amazonaws.com:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Create a gRPC client
	client := pb.NewListingServiceClient(conn)

	// Contact the server and print out its response.
	response, err := client.GetListingDetail(context.Background(), &pb.GetListingDetailRequest{ListingId: "listingId"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", response.ListingDetails)

	client2 := pb2.NewQueryServiceClient(conn)

	// Contact the server and print out its response.
	keys := []string{"GOOGLE_MAPS_KEY", "Yaobin"}
	response2, err2 := client2.GetValues(context.Background(), &pb2.GetValuesRequest{Keys: keys})
	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}
	log.Printf("Response: %v", response2.KeyValuePairs)
}
