package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type StoreProductRepository struct {
	pool *pgxpool.Pool
}

func NewStoreProductRepository(pool *pgxpool.Pool) *StoreProductRepository {
	return &StoreProductRepository{pool: pool}
}

