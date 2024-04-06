package main

import (
	"bufio"
	"keyfi-backend/apis/chat/ai"
	"keyfi-backend/apis/query"
	ai_pb "keyfi-backend/protos/ai"
	query_pb "keyfi-backend/protos/query"
	"keyfi-backend/util/chat"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"google.golang.org/grpc"
)

const CONFIG_MAPPINGS_PATH = "./config_mappings.key"

func listenOnGrpc() {
	// set a port for the server
	port := ":50051"

	// listen for requests on port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	grpcServer := grpc.NewServer()

	// Register different services
	ai_pb.RegisterAIServiceServer(grpcServer, &ai.Server{})
	query_pb.RegisterQueryServiceServer(grpcServer, &query.Server{})

	log.Printf("starting server on port %s\n", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}

func listenOnWebSocket() {
	http.HandleFunc("/", chat.HandleWebSocket)
	log.Print("starting websocket on port 50052")
	log.Fatal(http.ListenAndServe("0.0.0.0:50052", nil))
}

func main() {
	// Read Configs
	file, err := os.Open(CONFIG_MAPPINGS_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		os.Setenv(words[0], words[1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	go listenOnGrpc()
	listenOnWebSocket()
}
