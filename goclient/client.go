package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

// rpc client

type Args struct {
	Message string
}

func main() {

	argsWithoutProg := os.Args[1:]

	hostname := "localhost"
	port := ":1122"

	var reply string

	prompt := "hello, what is your name?"
	if len(argsWithoutProg) > 0 {
		prompt = argsWithoutProg[0]
	}

	fmt.Printf("prompt: %s\n", prompt)
	args := Args{prompt}

	client, err := rpc.DialHTTP("tcp", hostname+port)
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	// Call normally takes service name.function name, args and
	// the address of the variable that hold the reply. Here we
	// have no args in the demo therefore we can pass the empty
	// args struct.
	err = client.Call("BannerMessageServer.GetBannerMessage", args, &reply)
	if err != nil {
		log.Fatal("error", err)
	}

	// log the result
	log.Printf("response: %s\n", reply)
}
