package main

import (
	"log"
	"net"

	"github.com/rezaindrag/golang-grpc/example_2/common/config"

	"github.com/rezaindrag/golang-grpc/example_2/common/model"
	"google.golang.org/grpc"
)

var localStorage *model.GarageListByUser

func init() {
	localStorage = new(model.GarageListByUser)
	localStorage.List = make(map[string]*model.GarageList)
}

func main() {
	srv := grpc.NewServer()
	var garageSrv GaragesServer
	model.RegisterGaragesServer(srv, garageSrv)

	log.Println("Starting RPC server at", config.ServiceGaragePort)

	l, err := net.Listen("tcp", config.ServiceGaragePort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceGaragePort, err)
	}

	log.Fatal(srv.Serve(l))
}
