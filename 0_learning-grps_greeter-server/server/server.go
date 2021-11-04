package main

import (

	"log"
	"net"
	"context"
	"google.golang.org/grpc"
	"proto-app/greeter"
)

const (
	port = ":4000"
)

// server is used to implement greeter.TranslaterServer
type server struct {

	 greeter.UnimplementedTranslaterServer
}

// Dictionary implements greeter.TranslaterServer
func (s *server) Dictionary(ctx context.Context, in *greeter.Request) (*greeter.Response, error) {

	if in.GetData() == "Assalmu Alaikum" || in.GetData() == "salom" {

		return &greeter.Response { Data: "Vaalaikumussalam!"}, nil
	} 

	return &greeter.Response { Data: "Assalmu Alaikum!"}, nil 	
}

func main() {

	listen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greeter.RegisterTranslaterServer(s, &server{})

	err = s.Serve(listen)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}