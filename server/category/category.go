package category

import (
	"context"
	"fmt"
	"log"

	protoCategory "grpc-category/proto/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {}

type Category struct {
	ID	int64
	CategoryName string
	Status	bool
}

var categories = []Category{
	{
		ID : 1,
		CategoryName: "This is category 1",
		Status: true,
	},
	{
		ID : 2,
		CategoryName: "This is category 2",
		Status: true,
	},
	{
		ID : 3,
		CategoryName: "This is category 3",
		Status: true,
	},
	{
		ID : 4,
		CategoryName: "This is category 4",
		Status: true,
	},
}

func (s *Server) GetCategory(ctx context.Context, req *protoCategory.GetCategoryRequest) (*protoCategory.GetCategoryResponse, error) {
	log.Printf("Category ID: %d", req.GetID())
	var category Category
	for _, value := range categories {
		fmt.Println(value)
		if value.ID == req.GetID() {
			category = value
			break
		}
	}

	if category.ID == 0 {
		return &protoCategory.GetCategoryResponse{}, status.Errorf(codes.NotFound, "Invalid ID")
	}
	return &protoCategory.GetCategoryResponse{
		ID: category.ID,
		CategoryName: category.CategoryName,
		Status: category.Status,
	}, nil
}