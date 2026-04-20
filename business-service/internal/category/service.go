package category

import "context"

type repository interface {
	Create(ctx context.Context, req CreateRequest) (*Category, error)
	Update(ctx context.Context, category Category) (*Category, error)
	Delete(ctx context.Context, number int) error
	GetAll(ctx context.Context) ([]Category, error)
	GetByNumber(ctx context.Context, number int) (*Category, error)
	GetStockSummary(ctx context.Context) ([]StockSummary, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*Category, error) {
	return s.repo.Create(ctx, req)
}

func (s *Service) Update(ctx context.Context, category Category) (*Category, error) {
	return s.repo.Update(ctx, category)
}

func (s *Service) Delete(ctx context.Context, number int) error {
	return s.repo.Delete(ctx, number)
}

func (s *Service) GetAll(ctx context.Context) ([]Category, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetByNumber(ctx context.Context, number int) (*Category, error) {
	return s.repo.GetByNumber(ctx, number)
}

func (s *Service) GetStockSummary(ctx context.Context) ([]StockSummary, error) {
	return s.repo.GetStockSummary(ctx)
}
