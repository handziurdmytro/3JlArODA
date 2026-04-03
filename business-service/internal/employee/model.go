package employee

import "time"

type Employee struct {
	IdEmployee     string    `json:"id_employee"`
	EmplSurname    string    `json:"empl_surname"`
	EmplName       string    `json:"empl_name"`
	EmplPatronymic *string   `json:"empl_patronymic,omitempty"`
	EmplRole       string    `json:"empl_role"`
	Salary         float64   `json:"salary"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	DateOfStart    time.Time `json:"date_of_start"`
	PhoneNumber    string    `json:"phone_number"`
	City           string    `json:"city"`
	Street         string    `json:"street"`
	ZipCode        string    `json:"zip_code"`
}

type CreateRequest struct {
	EmplSurname    string    `json:"empl_surname"`
	EmplName       string    `json:"empl_name"`
	EmplPatronymic *string   `json:"empl_patronymic,omitempty"`
	EmplRole       string    `json:"empl_role"`
	Salary         float64   `json:"salary"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	DateOfStart    time.Time `json:"date_of_start"`
	PhoneNumber    string    `json:"phone_number"`
	City           string    `json:"city"`
	Street         string    `json:"street"`
	ZipCode        string    `json:"zip_code"`
}
