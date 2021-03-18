package main

import (
	"context"
	"log"

	"github.com/rezaindrag/golang-grpc/example_2/common/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UsersServer struct {
	model.UnimplementedUsersServer
}

func (UsersServer) Register(ctx context.Context, param *model.User) (*emptypb.Empty, error) {
	localStorage.List = append(localStorage.List, param)

	log.Println("Registering user", param.String())

	return new(emptypb.Empty), nil
}

func (UsersServer) List(ctx context.Context, void *emptypb.Empty) (*model.UserList, error) {
	return localStorage, nil
}
