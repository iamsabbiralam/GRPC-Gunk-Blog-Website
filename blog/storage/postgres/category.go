package postgres

import (
	"context"
	"grpc-blog/blog/storage"
)

const insertCategory = `INSERT INTO categories (category_name) VALUES (:category_name);`

func (s *Storage) Create(ctx context.Context, t storage.Category) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertCategory)
	if err != nil {
		return 0, nil
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, nil
	}
	return id, nil
}