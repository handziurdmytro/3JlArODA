package employee

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	pool *pgxpool.Pool
}

func NewEmployeeRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (repository Repository) CreateEmployee() {

}
