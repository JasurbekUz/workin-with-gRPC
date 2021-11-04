package main

import (

	"log"
	"os"
	"context"
	"proto-app/greeter"
	"google.golang.org/grpc"
)

const (

	address = "localhost:4000"
	defaultWord = "salomatmisiz?"
)

func main () {

	// Set up a connection to the server.
	connection , err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		
		log.Fatalf("clent connection error: %v", err)
	}

	defer connection.Close()

	stub := greeter.NewTranslaterClient(connection)

	// Contact the server and print out its response.
	word := defaultWord

	if len(os.Args) > 1 {

		word = os.Args[1]
	}

	res, err := stub.Dictionary(context.Background(), &greeter.Request{Data: word})

	if err != nil {
		
		log.Fatalf("%v", err)
	}

	log.Println(res)
}

