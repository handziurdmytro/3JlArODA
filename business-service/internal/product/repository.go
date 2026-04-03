package product

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}
