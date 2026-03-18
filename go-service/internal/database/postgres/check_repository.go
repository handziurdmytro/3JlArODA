package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type CheckRepository struct {
	pool *pgxpool.Pool
}

func NewCheckRepository(pool *pgxpool.Pool) *CheckRepository {
	return &CheckRepository{pool: pool}
}

