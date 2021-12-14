package postgres

import "github.com/jmoiron/sqlx"

type Storage struct {
	db	*sqlx.DB
}

func NewStorage() *Storage	{
	return &Storage{}
}