package category

import (
	"context"

	"grpc-blog/blog/storage"
)

type categoryStore interface {
	Create(context.Context, storage.Category) (int64, error)
}

type CoreSvc struct {
	store categoryStore
}

func NewCoreSvc(s categoryStore) *CoreSvc {
	return &CoreSvc{
		store: s,
	}
}

func (cs *CoreSvc) Create(ctx context.Context, c storage.Category) (int64, error) {
	return 0, nil
}