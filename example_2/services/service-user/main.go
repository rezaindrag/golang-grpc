package main

import (
	"log"
	"net"

	"github.com/rezaindrag/golang-grpc/example_2/common/config"

	"github.com/rezaindrag/golang-grpc/example_2/common/model"
	"google.golang.org/grpc"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

func main() {
	srv := grpc.NewServer()
	model.RegisterUsersServer(srv, &UsersServer{})

	log.Println("Starting RPC server at", config.ServiceUserPort)

	l, err := net.Listen("tcp", config.ServiceUserPort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceUserPort, err)
	}

	log.Fatal(srv.Serve(l))
}
