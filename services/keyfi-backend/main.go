package main

import (
	"bufio"
	"fmt"
	"keyfi-backend/apis/auth"
	"keyfi-backend/apis/chat/ai"
	"keyfi-backend/apis/listing"
	"keyfi-backend/apis/query"
	"keyfi-backend/chat"
	ai_pb "keyfi-backend/protos/ai"
	auth_pb "keyfi-backend/protos/auth"
	listing_pb "keyfi-backend/protos/listing"
	query_pb "keyfi-backend/protos/query"
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

	grpcServer := grpc.NewServer(
	// grpc.ChainUnaryInterceptor(
	// 	// Order matters e.g. tracing interceptor have to create span first for the later exemplars to work.
	// 	// middleware.UnaryInterceptor(),
	// ),
	)

	// Register different services

	ai_pb.RegisterAIServiceServer(grpcServer, &ai.Server{})
	query_pb.RegisterQueryServiceServer(grpcServer, &query.Server{})
	auth_pb.RegisterAuthenticationServiceServer(grpcServer, &auth.Server{})
	listing_pb.RegisterListingServiceServer(grpcServer, &listing.Server{})
	log.Printf("starting grpc server on port %s\n", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}

func listenOnWebSocket() {
	port := "50052"
	http.HandleFunc("/", chat.HandleWebSocket)
	log.Printf("starting websocket server on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
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
