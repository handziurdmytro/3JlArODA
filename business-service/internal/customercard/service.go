package customercard

import "context"

type repository interface {
	Create(ctx context.Context, req CreateRequest) (*CustomerCard, error)
	Update(ctx context.Context, card CustomerCard) (*CustomerCard, error)
	Delete(ctx context.Context, cardNumber string) error
	GetAll(ctx context.Context) ([]CustomerCard, error)
	GetByNumber(ctx context.Context, cardNumber string) (*CustomerCard, error)
	GetByPercent(ctx context.Context, percent int) ([]CustomerCard, error)
	SearchBySurname(ctx context.Context, surname string) ([]CustomerCard, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*CustomerCard, error) {
	return s.repo.Create(ctx, req)
}

func (s *Service) Update(ctx context.Context, card CustomerCard) (*CustomerCard, error) {
	return s.repo.Update(ctx, card)
}

func (s *Service) Delete(ctx context.Context, cardNumber string) error {
	return s.repo.Delete(ctx, cardNumber)
}

func (s *Service) GetAll(ctx context.Context) ([]CustomerCard, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetByNumber(ctx context.Context, cardNumber string) (*CustomerCard, error) {
	return s.repo.GetByNumber(ctx, cardNumber)
}

func (s *Service) GetByPercent(ctx context.Context, percent int) ([]CustomerCard, error) {
	return s.repo.GetByPercent(ctx, percent)
}

func (s *Service) SearchBySurname(ctx context.Context, surname string) ([]CustomerCard, error) {
	return s.repo.SearchBySurname(ctx, surname)
}
