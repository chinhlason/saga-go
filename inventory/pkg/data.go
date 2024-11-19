package pkg

import (
	"context"
	"database/sql"
)

type Product struct {
	ID       int    `json:"id"`
	Status   string `json:"status"`
	Number   int    `json:"number"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type IRepository interface {
	Get(ctx context.Context) (int, error)
	Update(ctx context.Context, id int, number int) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &Repository{db: db}
}

func (r *Repository) Get(ctx context.Context) (int, error) {
	row := r.db.QueryRowContext(ctx, "SELECT number FROM inventory WHERE id = 1")
	var number int
	err := row.Scan(&number)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func (r *Repository) Update(ctx context.Context, id int, number int) error {
	_, err := r.db.ExecContext(ctx, "UPDATE inventory SET number = ? WHERE id = ?", number, id)
	return err
}
