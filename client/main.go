package main

import (
	"context"
	"log"

	"grpc-category/proto/category"

	"google.golang.org/grpc"
)

func main () {
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %s", err)
	}

	c := category.NewCategoryServiceClient(conn)

	res, err := c.GetCategory(context.Background(), &category.GetCategoryRequest{
		ID : 2,
	})

	if err != nil {
		log.Fatalf("Error while calling Get Category: %s", err)
	}

	log.Printf("Response ID: %#v", res)
}