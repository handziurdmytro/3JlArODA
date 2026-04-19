package check

import (
	"context"
	"time"
)

type repository interface {
	Create(ctx context.Context, req CreateRequest) error
	DeleteByNumber(ctx context.Context, number string) error
	GetAll(ctx context.Context) ([]Check, error)
	GetAllOfTheDayByCashier(ctx context.Context, employeeID string, day time.Time) ([]Check, error)
	GetAllOfThePeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) ([]Check, error)
	GetFullDataByNumber(ctx context.Context, number string) ([]FullCheckData, error)
	GetDetailsOfThePeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) ([]Detail, error)
	GetDetailsOfThePeriod(ctx context.Context, from, to time.Time) ([]Detail, error)
	GetSumOfThePeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) (float64, error)
	GetSumOfThePeriod(ctx context.Context, from, to time.Time) (float64, error)
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

func (s *Service) DeleteByNumber(ctx context.Context, number string) error {
	return s.repo.DeleteByNumber(ctx, number)
}

func (s *Service) GetAll(ctx context.Context) ([]Check, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetAllOfTheDayByCashier(ctx context.Context, employeeID string, day time.Time) ([]Check, error) {
	return s.repo.GetAllOfTheDayByCashier(ctx, employeeID, day)
}

func (s *Service) GetAllOfThePeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) ([]Check, error) {
	return s.repo.GetAllOfThePeriodByCashier(ctx, employeeID, from, to)
}

func (s *Service) GetFullDataByNumber(ctx context.Context, number string) ([]FullCheckData, error) {
	return s.repo.GetFullDataByNumber(ctx, number)
}

func (s *Service) GetDetailsOfThePeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) ([]Detail, error) {
	return s.repo.GetDetailsOfThePeriodByCashier(ctx, employeeID, from, to)
}

func (s *Service) GetDetailsOfThePeriod(ctx context.Context, from, to time.Time) ([]Detail, error) {
	return s.repo.GetDetailsOfThePeriod(ctx, from, to)
}

func (s *Service) GetSumOfThePeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) (float64, error) {
	return s.repo.GetSumOfThePeriodByCashier(ctx, employeeID, from, to)
}

func (s *Service) GetSumOfThePeriod(ctx context.Context, from, to time.Time) (float64, error) {
	return s.repo.GetSumOfThePeriod(ctx, from, to)
}
