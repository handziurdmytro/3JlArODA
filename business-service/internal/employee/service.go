package employee

import "context"

type repository interface {
	CreateEmployee(ctx context.Context, req CreateRequest) error
	UpdateEmployeeByID(ctx context.Context, employee Employee) error
	DeleteEmployeeByID(ctx context.Context, id string) error
	GetAllEmployees(ctx context.Context) ([]Employee, error)
	GetEmployeeByID(ctx context.Context, id string) (*Employee, error)
	GetEmployeesByRole(ctx context.Context, role string) ([]Employee, error)
	GetEmployeeDataBySurname(ctx context.Context, surname string) ([]ContactData, error)
	GetEmployeeDataByFullName(ctx context.Context, surname, name, patronymic string) ([]ContactData, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*Employee, error) {
	if err := s.repo.CreateEmployee(ctx, req); err != nil {
		return nil, err
	}

	return s.repo.GetEmployeeByID(ctx, req.ID)
}

func (s *Service) Update(ctx context.Context, employee Employee) (*Employee, error) {
	if err := s.repo.UpdateEmployeeByID(ctx, employee); err != nil {
		return nil, err
	}

	return s.repo.GetEmployeeByID(ctx, employee.ID)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.DeleteEmployeeByID(ctx, id)
}

func (s *Service) GetByID(ctx context.Context, id string) (*Employee, error) {
	return s.repo.GetEmployeeByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]Employee, error) {
	return s.repo.GetAllEmployees(ctx)
}

func (s *Service) GetByRole(ctx context.Context, role string) ([]Employee, error) {
	return s.repo.GetEmployeesByRole(ctx, role)
}

func (s *Service) GetContactsBySurname(ctx context.Context, surname string) ([]ContactData, error) {
	return s.repo.GetEmployeeDataBySurname(ctx, surname)
}

func (s *Service) GetContactsByFullName(ctx context.Context, surname, name, patronymic string) ([]ContactData, error) {
	return s.repo.GetEmployeeDataByFullName(ctx, surname, name, patronymic)
}
