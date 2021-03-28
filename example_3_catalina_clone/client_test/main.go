package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/jsonpb"

	"github.com/rezaindrag/golang-grpc/example_3_catalina_clone/model"

	"google.golang.org/grpc"
)

const serverPort = ":3000"

func newCatalogClient() (model.CatalogServiceClient, func()) {
	conn, err := grpc.Dial(serverPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("could not connect to", serverPort, err)
	}

	return model.NewCatalogServiceClient(conn), func() { conn.Close() }
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	catalogClient, closeCatalogConn := newCatalogClient()
	defer closeCatalogConn()

	catalog, err := catalogClient.GetByID(ctx, &model.CatalogId{Id: "catalog-id"})
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	marshaler := jsonpb.Marshaler{
		EmitDefaults: true,
	}
	if err := marshaler.Marshal(&buf, catalog); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())
}
