package main

import (
	"chat/interfaces"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

const GOOGLE_AI_API_PATH = "../.googleai_apikey"

// an RPC server in Go

type Args struct {
	Message string
}

type BannerMessageServer string

func (t *BannerMessageServer) GetBannerMessage(args *Args, reply *string) error {
	fmt.Printf("prompt: %s\n", args.Message)

	response := interfaces.SendTextPrompt(args.Message)
	partString := ""
	for i := range response.Candidates[0].Content.Parts {
		partString = fmt.Sprintf("%s%s", partString, response.Candidates[0].Content.Parts[i])
	}
	fmt.Printf("response: %s\n", partString)
	*reply = partString
	return nil
}

func main() {

	// Read API key for Google AI
	dat, err := os.ReadFile(GOOGLE_AI_API_PATH)
	if err != nil {
		panic(err)
	}
	os.Setenv("GOOGLEAI_API_KEY", string(dat))

	// create and register the rpc
	banner := new(BannerMessageServer)
	rpc.Register(banner)
	rpc.HandleHTTP()

	// set a port for the server
	port := ":1122"

	// listen for requests on 1122
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	http.Serve(listener, nil)
}
