package main

import (

	"net"
	"log"
	"context"
	"strconv"
	"calc_app/calc"
	"google.golang.org/grpc"

)

const PORT = ":4000"

type server struct {

	calc.UnimplementedCalculatorServer
}

func (s *server) Count (ctx context.Context, in *calc.Request) (*calc.Response, error) {

	return &calc.Response {Result: Counter(in.GetExpression())}, nil
}

func main () {

	listen, err := net.Listen("tcp", PORT)

	if err != nil {
		
		log.Fatalf("server listenning error: %v", err)
	}

	newServer := grpc.NewServer()

	calc.RegisterCalculatorServer(newServer, &server{})

	log.Printf("server is ready in port %v", PORT)

	err = newServer.Serve(listen)

	if err != nil {
		log.Fatalf("NewServer serving error: %v", err)
	}
}

func Counter (exp string) float32 {

	var snum1, snum2, opr string
	var num1, num2 int
	var stop int

	for i, s := range exp {

		if 41 < s && s < 48 {

			stop = i
			break
		}
	}

	for index, symbol := range exp {

		if 47 < symbol && symbol < 58 {

			if index < stop {

				snum1 += string(symbol)

			}

			if index > stop {

				snum2 += string(symbol)
			}	
		}

		if 48 > symbol {

			opr = string(symbol)
			stop = index
		}
	}

	num1, _ = strconv.Atoi(snum1)
	num2, _ = strconv.Atoi(snum2)

	if opr == "+" {
		return float32(num1 + num2)
	
	} else if opr == "-" {
		return float32(num1 - num2)

	} else if opr == "*" || opr == "." {
		return float32(num1 * num2)

	} else if opr == "/" {
		return float32(num1 / num2)
	}

	return 0
}