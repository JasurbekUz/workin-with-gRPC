package main

import(
	"os"
	"log"
	"context"
	"dict_app/dict"
	"google.golang.org/grpc"
)

const (
	SERVER_ADDRESS = "localhost:4000"
	defaultWord = "Lug'at"
)	

func main () {

	conn, err := grpc.Dial(SERVER_ADDRESS, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("client connection error: %v", err)
	}

	defer conn.Close()

	c := dict.NewTranslaterClient(conn)

	word := defaultWord

	if len(os.Args) > 1 {

		word = os.Args[1]
	}


	res, err := c.Dictionary(context.Background(), &dict.Request { Word: word})

	if err != nil {
		log.Fatalf("getting result error: %v", err)
	}

	log.Printf("%v - %v", word, res.Word) 
}