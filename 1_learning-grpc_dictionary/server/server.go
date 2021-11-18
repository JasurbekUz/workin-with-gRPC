package main
// rpc - remote procedure call
import (
	"log"
	"net"
	"context"
	"google.golang.org/grpc"
	"dict_app/dict"

)

const PORT = ":4000"

type server struct {

	 dict.UnimplementedTranslaterServer
}

func (s *server) Dictionary (ctx context.Context, r *dict.Request) (*dict. Response, error) {

	return &dict.Response { Word: Words(r.GetWord())}, nil
}

func main() {

	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		
		log.Fatalf("listenning error: %v", err)
	}

	
	s := grpc.NewServer() // grpc server yasaldi

	dict.RegisterTranslaterServer(s, &server{}) // yasalgan serverga tarjimon serveri o'rnatib qo'yildi

	log.Printf("server is ready in port %v", PORT)

	err = s.Serve(lis) // shu server tcp connectionda ishga tihirildi 

	if err != nil {
		log.Fatalf("NewServer serving error: %v")
	}
}

func Words( word string) string {

	words := map[string]string{
		"Lug'at":"Dictionary",
		"salom":"hello",
		"Xayr":"bye",
		"kitob":"book",
		"daftar":"notebook",
		"ruchka":"pen",
		"qalam":"pencil",
		"kompyuter":"computer",
		"tarmoq":"network",
		"ism":"name",
		"o'yin":"game",
		"rang":"color",
		"oq":"white",
		"qora":"black",
		"so'z":"word",
	}

	value, ok := words[word]

	if ok {

		return value
	}

	return "word is not fount \n so'z topilmadi"
}