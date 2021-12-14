package storage

type Category struct {
	ID				int64 `db:"id"`
	CategoryName	string `db:"category_name"`
	IsCompleted		bool `db:"is_completed"`
}