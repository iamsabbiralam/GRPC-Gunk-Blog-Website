package postgres

import (
	"context"
	"grpc-blog/blog/storage"
)

const insertCategory = `
	INSERT INTO categories (
			category_name
		) VALUES (
			:category_name
		) RETURNING id;
`

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

func (s *Storage) Show(ctx context.Context, c storage.Category) (storage.Category, error) {
	var t storage.Category
	if err := s.db.Select(&t, "SELECT * FROM categories", c); err != nil {
		return t, err
	}
	return t, nil
}

func (s *Storage) Get(ctx context.Context, id int64) (*storage.Category, error) {
	var t storage.Category
	if err := s.db.Get(&t, "SELECT * FROM categories WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &t, nil
}

const updateCategory = `
	UPDATE categories
	SET
		category_name = :category_name
	WHERE
		id = :id
	RETURNING *;
`

func (s *Storage) Update(ctx context.Context, t storage.Category) (*storage.Category, error) {
	stmt, err := s.db.PrepareNamed(updateCategory)
	if err != nil {
		return nil, err
	}
	var st storage.Category
	if err := stmt.Get(&st, t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Storage) Delete(ctx context.Context, id int64) error {
	var t storage.Category
	if err := s.db.Get(&t, "DELETE FROM categories WHERE id = $1 RETURNING *", id); err != nil {
		return err
	}
	return nil
}