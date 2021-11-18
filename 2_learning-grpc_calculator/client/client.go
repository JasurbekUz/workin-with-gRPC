package main

import (
	"os"
	"log"
	"context"
	"calc_app/calc"
	"google.golang.org/grpc"
)

const (

	SERVER_ADDRESS = "localhost:4000"
	default_expression = "5*5"
)

func main () {

	connection, err := grpc.Dial(SERVER_ADDRESS, grpc.WithInsecure(), grpc.WithBlock())

	defer connection.Close()

	if err != nil {
		log.Fatalf("connection error: %v", err)
	}

	newClinet := calc.NewCalculatorClient(connection)

	expression := default_expression

	if len(os.Args) > 1 {

		expression = os.Args[1]
	}

	result, err := newClinet.Count(context.Background(), &calc.Request{ Expression: expression})

	if err != nil {
		log.Fatalf("getting result error: %v", err)
	}

	log.Printf("%v = %v", expression, result.Result)
}