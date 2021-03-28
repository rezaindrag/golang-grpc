package main

import (
	"log"

	"github.com/rezaindrag/golang-grpc/example_3_catalina_clone/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
