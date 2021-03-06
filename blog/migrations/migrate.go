package main

import (
	"grpc-blog/blog/storage/postgres"
	"log"
)

func main() {
	if err := postgres.Migrate(); err != nil {
		log.Fatal(err)
	}
}