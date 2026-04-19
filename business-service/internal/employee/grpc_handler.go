package employee

import (
	"context"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	employeepb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/employee"
)

type GRPCHandler struct {
	employeepb.UnimplementedEmployeeServiceServer

	service *Service
}

func NewGRPCHandler(service *Service) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (h *GRPCHandler) CreateEmployee(ctx context.Context, req *employeepb.CreateEmployeeRequest) (*employeepb.EmployeeResponse, error) {
	dateOfBirth, err := common.ParseProtoTime(req.GetDateOfBirth())
	if err != nil {
		return nil, err
	}

	dateOfStart, err := common.ParseProtoTime(req.GetDateOfStart())
	if err != nil {
		return nil, err
	}

	employee, err := h.service.Create(ctx, CreateRequest{
		ID:          req.GetId(),
		Surname:     req.GetSurname(),
		Name:        req.GetName(),
		Patronymic:  req.Patronymic,
		Role:        req.GetRole(),
		Salary:      req.GetSalary(),
		DateOfBirth: dateOfBirth,
		DateOfStart: dateOfStart,
		PhoneNumber: req.GetPhoneNumber(),
		City:        req.GetCity(),
		Street:      req.GetStreet(),
		ZipCode:     req.GetZipCode(),
	})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &employeepb.EmployeeResponse{Employee: toProtoEmployee(employee)}, nil
}

func (h *GRPCHandler) UpdateEmployee(ctx context.Context, req *employeepb.UpdateEmployeeRequest) (*employeepb.EmployeeResponse, error) {
	dateOfBirth, err := common.ParseProtoTime(req.GetDateOfBirth())
	if err != nil {
		return nil, err
	}

	dateOfStart, err := common.ParseProtoTime(req.GetDateOfStart())
	if err != nil {
		return nil, err
	}

	employee, err := h.service.Update(ctx, Employee{
		ID:          req.GetId(),
		Surname:     req.GetSurname(),
		Name:        req.GetName(),
		Patronymic:  req.Patronymic,
		Role:        req.GetRole(),
		Salary:      req.GetSalary(),
		DateOfBirth: dateOfBirth,
		DateOfStart: dateOfStart,
		PhoneNumber: req.GetPhoneNumber(),
		City:        req.GetCity(),
		Street:      req.GetStreet(),
		ZipCode:     req.GetZipCode(),
	})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &employeepb.EmployeeResponse{Employee: toProtoEmployee(employee)}, nil
}

func (h *GRPCHandler) DeleteEmployee(ctx context.Context, req *employeepb.DeleteEmployeeRequest) (*employeepb.DeleteEmployeeResponse, error) {
	if err := h.service.Delete(ctx, req.GetId()); err != nil {
		return nil, common.ToStatusError(err)
	}

	return &employeepb.DeleteEmployeeResponse{Success: true}, nil
}

func (h *GRPCHandler) GetEmployee(ctx context.Context, req *employeepb.GetEmployeeRequest) (*employeepb.EmployeeResponse, error) {
	employee, err := h.service.GetByID(ctx, req.GetId())
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &employeepb.EmployeeResponse{Employee: toProtoEmployee(employee)}, nil
}

func (h *GRPCHandler) GetAllEmployees(ctx context.Context, _ *employeepb.GetAllEmployeesRequest) (*employeepb.GetEmployeesResponse, error) {
	employees, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &employeepb.GetEmployeesResponse{Employees: toProtoEmployees(employees)}, nil
}

func (h *GRPCHandler) GetEmployeesByRole(ctx context.Context, req *employeepb.GetEmployeesByRoleRequest) (*employeepb.GetEmployeesResponse, error) {
	employees, err := h.service.GetByRole(ctx, req.GetRole())
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &employeepb.GetEmployeesResponse{Employees: toProtoEmployees(employees)}, nil
}

func (h *GRPCHandler) GetEmployeeContactsBySurname(ctx context.Context, req *employeepb.GetEmployeeContactsBySurnameRequest) (*employeepb.GetEmployeeContactsResponse, error) {
	contacts, err := h.service.GetContactsBySurname(ctx, req.GetSurname())
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &employeepb.GetEmployeeContactsResponse{Contacts: toProtoContacts(contacts)}, nil
}

func (h *GRPCHandler) GetEmployeeContactsByFullName(ctx context.Context, req *employeepb.GetEmployeeContactsByFullNameRequest) (*employeepb.GetEmployeeContactsResponse, error) {
	contacts, err := h.service.GetContactsByFullName(ctx, req.GetSurname(), req.GetName(), req.GetPatronymic())
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &employeepb.GetEmployeeContactsResponse{Contacts: toProtoContacts(contacts)}, nil
}

func toProtoEmployees(employees []Employee) []*employeepb.Employee {
	result := make([]*employeepb.Employee, 0, len(employees))
	for _, employee := range employees {
		employee := employee
		result = append(result, toProtoEmployee(&employee))
	}

	return result
}

func toProtoEmployee(employee *Employee) *employeepb.Employee {
	if employee == nil {
		return nil
	}

	return &employeepb.Employee{
		Id:          employee.ID,
		Surname:     employee.Surname,
		Name:        employee.Name,
		Patronymic:  employee.Patronymic,
		Role:        employee.Role,
		Salary:      employee.Salary,
		DateOfBirth: common.FormatProtoTime(employee.DateOfBirth),
		DateOfStart: common.FormatProtoTime(employee.DateOfStart),
		PhoneNumber: employee.PhoneNumber,
		City:        employee.City,
		Street:      employee.Street,
		ZipCode:     employee.ZipCode,
	}
}

func toProtoContacts(contacts []ContactData) []*employeepb.EmployeeContact {
	result := make([]*employeepb.EmployeeContact, 0, len(contacts))
	for _, contact := range contacts {
		result = append(result, &employeepb.EmployeeContact{
			PhoneNumber: contact.PhoneNumber,
			City:        contact.City,
			Street:      contact.Street,
			ZipCode:     contact.ZipCode,
		})
	}

	return result
}
