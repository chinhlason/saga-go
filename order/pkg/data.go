package pkg

import (
	"context"
	"database/sql"
)

type Order struct {
	ID       int    `json:"id"`
	Status   string `json:"status"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type IRepository interface {
	Insert(ctx context.Context) error
	Get(ctx context.Context, id int) (*Order, error)
	Update(ctx context.Context, id int, status string) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &Repository{db: db}
}

func (r *Repository) Insert(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO orders (status) VALUES ($1)", "PENDING")
	return err
}

func (r *Repository) Get(ctx context.Context, id int) (*Order, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM orders WHERE id = ?", id)
	order := &Order{}
	err := row.Scan(&order.ID, &order.Status, &order.CreateAt, &order.UpdateAt)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *Repository) Update(ctx context.Context, id int, status string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE orders SET status = ? WHERE id = ?", status, id)
	return err
}
