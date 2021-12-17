package category

import (
	"context"

	"grpc-blog/blog/storage"
)

type categoryStore interface {
	Create(context.Context, storage.Category) (int64, error)
	Get(ctx context.Context, t storage.Category) (storage.Category, error)
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
	return cs.store.Create(ctx, c)
}

func (cs *CoreSvc) Get(ctx context.Context, t storage.Category) (storage.Category, error) {
	return cs.store.Get(ctx, t)
}