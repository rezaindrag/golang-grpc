package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/rezaindrag/golang-grpc/example_2/common/config"
	"github.com/rezaindrag/golang-grpc/example_2/common/model"
	"google.golang.org/grpc"
)

func serviceGarage() (model.GaragesClient, func()) {
	port := config.ServiceGaragePort
	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn), func() { conn.Close() }
}

func serviceUser() (model.UsersClient, func()) {
	port := config.ServiceUserPort
	conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn), func() { conn.Close() }
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userClient, closeUserConn := serviceUser()
	defer closeUserConn()

	garageClient, closeGarageConn := serviceGarage()
	defer closeGarageConn()

	// Register user
	user1 := model.User{
		Id:       "n001",
		Name:     "Noval Agung",
		Password: "kw8d hl12/3m,a",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}
	_, err := userClient.Register(ctx, &user1)
	if err != nil {
		log.Fatal("error registering user:", err.Error())
	}
	res1, err := userClient.List(ctx, new(emptypb.Empty))
	if err != nil {
		log.Fatal("error listing user:", err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	// Add garage
	garage1 := model.Garage{
		Id:   "q001",
		Name: "Quel'thalas",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}
	_, err = garageClient.Add(ctx, &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})
	if err != nil {
		log.Fatal("error adding garage:", err.Error())
	}
	res2, err := garageClient.List(ctx, &model.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal("error listing garage:", err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))
}
