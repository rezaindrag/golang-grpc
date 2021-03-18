package main

import (
	"context"
	"log"

	"github.com/rezaindrag/golang-grpc/example_2/common/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GaragesServer struct {
	model.UnimplementedGaragesServer
}

func (GaragesServer) Add(ctx context.Context, param *model.GarageAndUserId) (*emptypb.Empty, error) {
	userId := param.UserId
	garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(model.GarageList)
		localStorage.List[userId].List = make([]*model.Garage, 0)
	}
	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	log.Println("Adding garage", garage.String(), "for user", userId)

	return new(emptypb.Empty), nil
}

func (GaragesServer) List(ctx context.Context, param *model.GarageUserId) (*model.GarageList, error) {
	userId := param.UserId

	return localStorage.List[userId], nil
}
