package employee

import (
	"context"
	"fmt"
	"time"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	employeedb "github.com/handziurdmytro/3JlArODA/business-service/internal/employee/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	queries *employeedb.Queries
}

type ContactData struct {
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	Street      string `json:"street"`
	ZipCode     string `json:"zip_code"`
}

func NewEmployeeRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{queries: employeedb.New(pool)}
}

func (r *Repository) CreateEmployee(ctx context.Context, req CreateRequest) error {
	salary, err := common.NumericFromFloat64(req.Salary)
	if err != nil {
		return fmt.Errorf("create employee %q: %w", req.ID, err)
	}

	err = r.queries.CreateEmployee(ctx, employeedb.CreateEmployeeParams{
		IDEmployee:     req.ID,
		EmplSurname:    req.Surname,
		EmplName:       req.Name,
		EmplPatronymic: common.TextFromPtr(req.Patronymic),
		EmplRole:       req.Role,
		Salary:         salary,
		DateOfBirth:    common.DateFromTime(req.DateOfBirth),
		DateOfStart:    common.DateFromTime(req.DateOfStart),
		PhoneNumber:    req.PhoneNumber,
		City:           req.City,
		Street:         req.Street,
		ZipCode:        req.ZipCode,
	})
	if err != nil {
		return common.MapRepositoryError(fmt.Sprintf("create employee %q", req.ID), err)
	}

	return nil
}

func (r *Repository) UpdateEmployeeByID(ctx context.Context, employee Employee) error {
	if _, err := r.GetEmployeeByID(ctx, employee.ID); err != nil {
		return err
	}

	salary, err := common.NumericFromFloat64(employee.Salary)
	if err != nil {
		return fmt.Errorf("update employee %q: %w", employee.ID, err)
	}

	err = r.queries.UpdateEmployeeByID(ctx, employeedb.UpdateEmployeeByIDParams{
		IDEmployee:     employee.ID,
		EmplSurname:    employee.Surname,
		EmplName:       employee.Name,
		EmplPatronymic: common.TextFromPtr(employee.Patronymic),
		EmplRole:       employee.Role,
		Salary:         salary,
		DateOfBirth:    common.DateFromTime(employee.DateOfBirth),
		DateOfStart:    common.DateFromTime(employee.DateOfStart),
		PhoneNumber:    employee.PhoneNumber,
		City:           employee.City,
		Street:         employee.Street,
		ZipCode:        employee.ZipCode,
	})
	if err != nil {
		return common.MapRepositoryError(fmt.Sprintf("update employee %q", employee.ID), err)
	}

	return nil
}

func (r *Repository) DeleteEmployeeByID(ctx context.Context, id string) error {
	if _, err := r.GetEmployeeByID(ctx, id); err != nil {
		return err
	}

	if err := r.queries.DeleteEmployeeByID(ctx, id); err != nil {
		return common.MapRepositoryError(fmt.Sprintf("delete employee %q", id), err)
	}

	return nil
}

func (r *Repository) GetAllEmployees(ctx context.Context) ([]Employee, error) {
	rows, err := r.queries.GetAllEmployees(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get all employees", err)
	}

	return mapEmployees(rows)
}

func (r *Repository) GetEmployeeByID(ctx context.Context, id string) (*Employee, error) {
	row, err := r.queries.GetEmployeeByID(ctx, id)
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get employee by id %q", id), err)
	}

	return mapEmployee(row)
}

func (r *Repository) GetEmployeesByRole(ctx context.Context, role string) ([]Employee, error) {
	rows, err := r.queries.GetEmployeesByRole(ctx, role)
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get employees by role %q", role), err)
	}

	return mapEmployees(rows)
}

func (r *Repository) GetEmployeeDataBySurname(ctx context.Context, surname string) ([]ContactData, error) {
	rows, err := r.queries.GetEmployeeDataBySurname(ctx, surname)
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get employee contact data by surname %q", surname), err)
	}

	contacts := make([]ContactData, 0, len(rows))
	for _, row := range rows {
		contacts = append(contacts, ContactData{
			PhoneNumber: row.PhoneNumber,
			City:        row.City,
			Street:      row.Street,
			ZipCode:     row.ZipCode,
		})
	}

	return contacts, nil
}

func (r *Repository) GetEmployeeDataByFullName(ctx context.Context, surname, name, patronymic string) ([]ContactData, error) {
	rows, err := r.queries.GetEmployeeDataByFullName(ctx, employeedb.GetEmployeeDataByFullNameParams{
		EmplSurname:    surname,
		EmplName:       name,
		EmplPatronymic: common.TextFromString(patronymic),
	})
	if err != nil {
		return nil, common.MapRepositoryError(
			fmt.Sprintf("get employee contact data by full name %q %q %q", surname, name, patronymic),
			err,
		)
	}

	contacts := make([]ContactData, 0, len(rows))
	for _, row := range rows {
		contacts = append(contacts, ContactData{
			PhoneNumber: row.PhoneNumber,
			City:        row.City,
			Street:      row.Street,
			ZipCode:     row.ZipCode,
		})
	}

	return contacts, nil
}

func (r *Repository) GetCashierPerformance(ctx context.Context, from, to time.Time, minRevenue float64) ([]CashierPerformance, error) {
	minRevenueValue, err := common.NumericFromFloat64(minRevenue)
	if err != nil {
		return nil, fmt.Errorf("get cashier performance: %w", err)
	}

	rows, err := r.queries.GetCashierPerformance(ctx, employeedb.GetCashierPerformanceParams{
		PrintDate:    common.TimestampFromTime(from),
		PrintDate_2:  common.TimestampFromTime(to),
		SellingPrice: minRevenueValue,
	})
	if err != nil {
		return nil, common.MapRepositoryError("get cashier performance", err)
	}

	performance := make([]CashierPerformance, 0, len(rows))
	for _, row := range rows {
		totalRevenue, err := common.Float64FromNumeric(row.TotalRevenue)
		if err != nil {
			return nil, fmt.Errorf("map cashier performance total revenue: %w", err)
		}

		performance = append(performance, CashierPerformance{
			ID:             row.IDEmployee,
			Name:           row.EmplName,
			Surname:        row.EmplSurname,
			Patronymic:     common.PtrFromText(row.EmplPatronymic),
			TotalChecks:    row.TotalChecks,
			TotalItemsSold: row.TotalItemsSold,
			TotalRevenue:   totalRevenue,
		})
	}

	return performance, nil
}

func (r *Repository) GetBestCashiersByPromo(ctx context.Context) ([]BestCashierByPromo, error) {
	rows, err := r.queries.GetBestCashiersByPromo(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get best cashiers by promo", err)
	}

	cashiers := make([]BestCashierByPromo, 0, len(rows))
	for _, row := range rows {
		cashiers = append(cashiers, BestCashierByPromo{
			ID:      row.IDEmployee,
			Name:    row.EmplName,
			Surname: row.EmplSurname,
		})
	}

	return cashiers, nil
}

func mapEmployees(rows []employeedb.Employee) ([]Employee, error) {
	employees := make([]Employee, 0, len(rows))
	for _, row := range rows {
		employee, err := mapEmployee(row)
		if err != nil {
			return nil, err
		}
		employees = append(employees, *employee)
	}

	return employees, nil
}

func mapEmployee(row employeedb.Employee) (*Employee, error) {
	salary, err := common.Float64FromNumeric(row.Salary)
	if err != nil {
		return nil, fmt.Errorf("map employee salary: %w", err)
	}

	return &Employee{
		ID:          row.IDEmployee,
		Surname:     row.EmplSurname,
		Name:        row.EmplName,
		Patronymic:  common.PtrFromText(row.EmplPatronymic),
		Role:        row.EmplRole,
		Salary:      salary,
		DateOfBirth: common.TimeFromDate(row.DateOfBirth),
		DateOfStart: common.TimeFromDate(row.DateOfStart),
		PhoneNumber: row.PhoneNumber,
		City:        row.City,
		Street:      row.Street,
		ZipCode:     row.ZipCode,
	}, nil
}
