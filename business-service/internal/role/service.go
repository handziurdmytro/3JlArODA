package role

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetEmployeeRole(ctx context.Context, id string) (string, error) {
	return s.repo.GetEmployeeRole(ctx, id)
}
