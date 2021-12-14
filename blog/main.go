package main

import (
	"log"
	"net"

	coreCategory "grpc-blog/blog/core/category"
	"grpc-blog/blog/services/category"
	"grpc-blog/blog/storage/postgres"
	protoCategory "grpc-blog/gunk/v1/category"

	"google.golang.org/grpc"
)

func main () {
	lis, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}

	grpcServer := grpc.NewServer()
	store := postgres.NewStorage()
	cs := coreCategory.NewCoreSvc(store)
	s := category.NewCategoryServer(cs)

	protoCategory.RegisterCategoryServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed too serve: %s", err)
	}
}