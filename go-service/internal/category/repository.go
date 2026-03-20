package category

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	pool *pgxpool.Pool
}

func NewCategoryRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}
