package postgres

import (
	"context"
	"grpc-blog/blog/storage"
)

const insertCategory = `INSERT INTO categories (category_name) VALUES (:category_name) RETURNING id;`

func (s *Storage) Create(ctx context.Context, t storage.Category) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertCategory)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) Get(ctx context.Context, id int64) (*storage.Category, error) {
	var t storage.Category
	if err := s.db.Get(&t, "SELECT * FROM categories WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &t, nil
}