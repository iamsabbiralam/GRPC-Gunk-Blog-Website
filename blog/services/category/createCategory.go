package category

import (
	"context"

	"grpc-blog/blog/storage"
	protoCategory "grpc-blog/gunk/v1/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) Create(ctx context.Context, req *protoCategory.CreateCategoryRequest) (*protoCategory.CreateCategoryResponse, error) {
	// need to validate request
	category := storage.Category{
		ID: req.GetCategory().ID,
		CategoryName: req.Category.CategoryName,
	}
	id, err := s.core.Create(context.Background(), category)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create category")
	}
	return &protoCategory.CreateCategoryResponse{
		ID: id,
	}, nil
}