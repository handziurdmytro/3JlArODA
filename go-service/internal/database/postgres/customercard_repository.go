package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type CustomerCardRepository struct {
	pool *pgxpool.Pool
}

func NewCustomerCardRepository(pool *pgxpool.Pool) *CustomerCardRepository {
	return &CustomerCardRepository{pool: pool}
}

