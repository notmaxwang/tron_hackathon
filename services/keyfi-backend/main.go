package main

import (
	"bufio"
	"keyfi-backend/apis/chat/ai"
	pb "keyfi-backend/protos"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
)

const APIKEY_MAPPINGS_PATH = "./apikey_mappings.key"

func main() {
	// Read API key for Google AI
	file, err := os.Open(APIKEY_MAPPINGS_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		os.Setenv(words[0], words[1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// set a port for the server
	port := ":50051"

	// listen for requests on port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	grpcServer := grpc.NewServer()

	// Register different services
	pb.RegisterAIServiceServer(grpcServer, &ai.Server{})


	log.Printf("starting server on port %s\n", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
