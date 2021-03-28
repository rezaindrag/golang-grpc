package grpc_server

import (
	"context"
	"time"

	catalina "github.com/rezaindrag/golang-grpc/example_3_catalina_clone"

	"github.com/rezaindrag/golang-grpc/example_3_catalina_clone/model"
)

type catalogGRPCServer struct {
	model.UnimplementedCatalogServiceServer
}

func NewCatalogGRPCServer() model.CatalogServiceServer {
	return &catalogGRPCServer{}
}

func (g catalogGRPCServer) GetByID(context.Context, *model.CatalogId) (*model.Catalog, error) {
	catalog := catalina.Catalog{
		ID: "catalog-id",
		Items: []catalina.Promo{
			{
				ID: "promo-1",
			},
			{
				ID: "promo-2",
			},
		},
		StartTime: time.Now(),
		EndTime:   time.Now(),
	}

	return catalog.ConvertToGRPCStruct(), nil
}
