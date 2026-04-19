package sale

import (
	"context"
	"time"
)

type repository interface {
	Create(ctx context.Context, req CreateRequest) error
	GetProductSoldQuantity(ctx context.Context, productID int, from, to time.Time) (int64, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) error {
	return s.repo.Create(ctx, req)
}

func (s *Service) GetProductSoldQuantity(ctx context.Context, productID int, from, to time.Time) (int64, error) {
	return s.repo.GetProductSoldQuantity(ctx, productID, from, to)
}
