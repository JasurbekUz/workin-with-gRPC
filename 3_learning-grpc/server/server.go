package main

import (

	"log"
	"net"
	"context"

	"app/dbapp"
	d "app/dbconfigure"

	"google.golang.org/grpc"
)

const PORT = ":4000"

var user  dbapp.UserProfile

type server struct{

	dbapp.UnimplementedUsersInfoServer
}

func (s *server) PostUser(ctx context.Context, in *dbapp.PostUserReq) (*dbapp.UserProfile, error) {

	db := d.DataBase()

	var user  dbapp.UserProfile

	err := db.QueryRow("insert into users (full_name, user_name, phone_number) values ($1, $2, $3) returning full_name, user_name, phone_number", 
						in.GetNewUser().GetFullName(), 
						in.GetNewUser().GetUserName(), 
						in.GetNewUser().GetPhoneNumber(),
			).Scan(
				&user.FullName, 
				&user.UserName, 
				&user.PhoneNumber,
			)

	if err != nil {

		return &user, nil
	}

	return &user, nil

}

func (s *server) GetUser(ctx context.Context, in *dbapp.GetUserReq) (*dbapp.UserProfile, error) {

	var user  dbapp.UserProfile

	db := d.DataBase()

	err := db.QueryRow("select full_name, user_name, phone_number from users where user_id = $1", in.GetUserId()).Scan(&user.FullName, &user.UserName, &user.PhoneNumber)

	if err != nil {

		return &user, nil
	}

	return &user, nil

}

/*func (s *server) UpdateUser(ctx context.Context, in *dbapp.UpdateUserReq) (*dbapp.UserProfile, error) {

	db := d.DataBase()

	var user  dbapp.UserProfile

	err := db.QueryRow("update users set full_name = COALESCE($1, full_name), phone_number = COALESCE($2, phone_number) where user_id = 5 returning full_name, user_name, phone_number", 
						in.GetUpdatedUser().GetFullName(), 
						in.GetUpdatedUser().GetUserName(), 
						in.GetUpdatedUser().GetPhoneNumber(),
			).Scan(
				&user.FullName, 
				&user.UserName, 
				&user.PhoneNumber,
			)

	if err != nil {

		return &user, nil
	}

	return &user, nil

}*/

func (s *server) DeleteUser(ctx context.Context, in *dbapp.DeleteUserReq) (*dbapp.UserProfile, error) {

	var user  dbapp.UserProfile

	db := d.DataBase()

	err := db.QueryRow("delete from users where user_id = $1 returning full_name, user_name, phone_number", 
						in.GetUserId(),
					).Scan(
						&user.FullName, 
						&user.UserName, 
						&user.PhoneNumber,
					)

	if err != nil {

		return &user, nil
	}

	return &user, nil

}

func main () {

	listen, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("listenning error: %v", err)
	}

	newServer := grpc.NewServer()

	dbapp.RegisterUsersInfoServer(newServer, &server{})

	log.Printf("server is listenning at %v", listen.Addr())

	if err := newServer.Serve(listen); err != nil {
		
		log.Fatalf("server didn't return: %v", err)
	}
}