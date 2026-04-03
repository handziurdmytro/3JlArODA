package storeproduct

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	pool *pgxpool.Pool
}

func NewStoreProductRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}
