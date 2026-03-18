package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type SaleRepository struct {
	pool *pgxpool.Pool
}

func NewSaleRepository(pool *pgxpool.Pool) *SaleRepository {
	return &SaleRepository{pool: pool}
}

