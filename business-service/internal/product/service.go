package product

import "context"

type repository interface {
	Create(ctx context.Context, req CreateRequest) (*Product, error)
	Update(ctx context.Context, product Product) (*Product, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id int) (*Product, error)
	GetByCategory(ctx context.Context, categoryNumber int) ([]Product, error)
	SearchByName(ctx context.Context, name string) ([]Product, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*Product, error) {
	return s.repo.Create(ctx, req)
}

func (s *Service) Update(ctx context.Context, product Product) (*Product, error) {
	return s.repo.Update(ctx, product)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id int) (*Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetByCategory(ctx context.Context, categoryNumber int) ([]Product, error) {
	return s.repo.GetByCategory(ctx, categoryNumber)
}

func (s *Service) SearchByName(ctx context.Context, name string) ([]Product, error) {
	return s.repo.SearchByName(ctx, name)
}
