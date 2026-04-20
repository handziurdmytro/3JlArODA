package clients

import (
	"context"
	"time"

	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	employeepb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/employee"
	"google.golang.org/grpc"
)

type EmployeeClient struct {
	client employeepb.EmployeeServiceClient
}

func NewEmployeeClient(conn grpc.ClientConnInterface) *EmployeeClient {
	return &EmployeeClient{
		client: employeepb.NewEmployeeServiceClient(conn),
	}
}

func (cl *EmployeeClient) Create(ctx context.Context, req models.CreateEmployeeRequest) (*employeepb.Employee, error) {
	resp, err := cl.client.CreateEmployee(ctx, &employeepb.CreateEmployeeRequest{
		Id:          req.ID,
		Surname:     req.Surname,
		Name:        req.Name,
		Patronymic:  req.Patronymic,
		Role:        req.Role,
		Salary:      req.Salary,
		DateOfBirth: req.DateOfBirth.Format(timeFormat),
		DateOfStart: req.DateOfStart.Format(timeFormat),
		PhoneNumber: req.PhoneNumber,
		City:        req.City,
		Street:      req.Street,
		ZipCode:     req.ZipCode,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetEmployee(), nil
}

func (cl *EmployeeClient) Update(ctx context.Context, id string, req models.UpdateEmployeeRequest) (*employeepb.Employee, error) {
	resp, err := cl.client.UpdateEmployee(ctx, &employeepb.UpdateEmployeeRequest{
		Id:          id,
		Surname:     req.Surname,
		Name:        req.Name,
		Patronymic:  req.Patronymic,
		Role:        req.Role,
		Salary:      req.Salary,
		DateOfBirth: req.DateOfBirth.Format(timeFormat),
		DateOfStart: req.DateOfStart.Format(timeFormat),
		PhoneNumber: req.PhoneNumber,
		City:        req.City,
		Street:      req.Street,
		ZipCode:     req.ZipCode,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetEmployee(), nil
}

func (cl *EmployeeClient) Delete(ctx context.Context, id string) error {
	_, err := cl.client.DeleteEmployee(ctx, &employeepb.DeleteEmployeeRequest{Id: id})
	return err
}

func (cl *EmployeeClient) GetByID(ctx context.Context, id string) (*employeepb.Employee, error) {
	resp, err := cl.client.GetEmployee(ctx, &employeepb.GetEmployeeRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return resp.GetEmployee(), nil
}

func (cl *EmployeeClient) GetAll(ctx context.Context) ([]*employeepb.Employee, error) {
	resp, err := cl.client.GetAllEmployees(ctx, &employeepb.GetAllEmployeesRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetEmployees(), nil
}

func (cl *EmployeeClient) GetByRole(ctx context.Context, role string) ([]*employeepb.Employee, error) {
	resp, err := cl.client.GetEmployeesByRole(ctx, &employeepb.GetEmployeesByRoleRequest{Role: role})
	if err != nil {
		return nil, err
	}

	return resp.GetEmployees(), nil
}

func (cl *EmployeeClient) GetContactsBySurname(ctx context.Context, surname string) ([]*employeepb.EmployeeContact, error) {
	resp, err := cl.client.GetEmployeeContactsBySurname(ctx, &employeepb.GetEmployeeContactsBySurnameRequest{Surname: surname})
	if err != nil {
		return nil, err
	}

	return resp.GetContacts(), nil
}

func (cl *EmployeeClient) GetContactsByFullName(ctx context.Context, surname, name string, patronymic *string) ([]*employeepb.EmployeeContact, error) {
	resp, err := cl.client.GetEmployeeContactsByFullName(ctx, &employeepb.GetEmployeeContactsByFullNameRequest{
		Surname:    surname,
		Name:       name,
		Patronymic: patronymic,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetContacts(), nil
}

func (cl *EmployeeClient) GetCashierPerformance(ctx context.Context, from, to time.Time, minRevenue float64) ([]*employeepb.CashierPerformance, error) {
	resp, err := cl.client.GetCashierPerformance(ctx, &employeepb.GetCashierPerformanceRequest{
		From:       from.Format(timeFormat),
		To:         to.Format(timeFormat),
		MinRevenue: minRevenue,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetPerformance(), nil
}

func (cl *EmployeeClient) GetBestCashiersByPromo(ctx context.Context) ([]*employeepb.BestCashierByPromo, error) {
	resp, err := cl.client.GetBestCashiersByPromo(ctx, &employeepb.GetBestCashiersByPromoRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetCashiers(), nil
}
