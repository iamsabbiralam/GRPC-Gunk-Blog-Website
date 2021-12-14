package category

import (
	"context"

	"grpc-blog/blog/storage"
	protoCategory "grpc-blog/gunk/v1/category"
)

type categoryCoreStore interface {
	Create(context.Context, storage.Category) (int64, error)
}

type Svc struct {
	protoCategory.UnimplementedCategoryServiceServer
	core	categoryCoreStore
}

func NewCategoryServer(c categoryCoreStore) *Svc {
	return &Svc{
		core: c,
	}
}