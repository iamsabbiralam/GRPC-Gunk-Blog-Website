package main

import (
	"grpc-category/client/cat"
	"log"

	"google.golang.org/grpc"
)

func main () {
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %s", err)
	}

	categoryClient := cat.NewClient(conn)
	res, err := categoryClient.GetCategory(2)

	if err != nil {
		log.Fatalf("Error while calling Get Category: %s", err)
	}

	log.Printf("Response ID: %#v", res)
	if err := categoryClient.GetCategories(); err != nil {
		log.Fatalf("Error while calling Get Category: %s", err)
	}
}