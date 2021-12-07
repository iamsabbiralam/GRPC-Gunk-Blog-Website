package main

import (
	"log"
	"net"
	
	"grpc-category/server/category"
	"google.golang.org/grpc"
	protoCategory "grpc-category/proto/category"
)

func main () {
	lis, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}

	grpcServer := grpc.NewServer()
	s := category.Server{}

	protoCategory.RegisterCategoryServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed too serve: %s", err)
	}
}