package cmd

import (
	"log"
	"net"

	"github.com/rezaindrag/golang-grpc/example_3_catalina_clone/internal/grpc_server"

	"github.com/rezaindrag/golang-grpc/example_3_catalina_clone/model"

	"google.golang.org/grpc"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc_server",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		srv := grpc.NewServer()
		model.RegisterCatalogServiceServer(srv, grpc_server.NewCatalogGRPCServer())

		l, err := net.Listen("tcp", viper.GetString("port"))
		if err != nil {
			log.Fatalf("could not listen to %s: %v", viper.GetString("port"), err)
		}

		log.Println("start grpc server with port", viper.GetString("port"))

		log.Fatal(srv.Serve(l))
	},
}
