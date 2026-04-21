package role

import (
	"context"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/role/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
	q  *roledb.Queries
}

func NewRoleRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
		q:  roledb.New(db),
	}
}

func (r *Repository) GetEmployeeRole(ctx context.Context, id string) (string, error) {
	return r.q.GetEmployeeRole(ctx, id)
}
