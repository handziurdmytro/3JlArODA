package models

import "time"

// Employee Models
type CreateEmployeeRequest struct {
	ID          string    `json:"id" binding:"required"`
	Surname     string    `json:"surname" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Patronymic  *string   `json:"patronymic,omitempty"`
	Role        string    `json:"role" binding:"required"`
	Salary      float64   `json:"salary" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	DateOfStart time.Time `json:"date_of_start" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	City        string    `json:"city" binding:"required"`
	Street      string    `json:"street" binding:"required"`
	ZipCode     string    `json:"zip_code" binding:"required"`
}

type UpdateEmployeeRequest struct {
	Surname     *string    `json:"surname"`
	Name        *string    `json:"name"`
	Patronymic  *string    `json:"patronymic,omitempty"`
	Role        *string    `json:"role"`
	Salary      *float64   `json:"salary"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	DateOfStart *time.Time `json:"date_of_start"`
	PhoneNumber *string    `json:"phone_number"`
	City        *string    `json:"city"`
	Street      *string    `json:"street"`
	ZipCode     *string    `json:"zip_code"`
}

// Category Models
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name *string `json:"name"`
}

// Product Models
type CreateProductRequest struct {
	CategoryNumber  int     `json:"category_number" binding:"required"`
	Name            string  `json:"name" binding:"required"`
	Producer        *string `json:"producer,omitempty"`
	Characteristics string  `json:"characteristics" binding:"required"`
}

type UpdateProductRequest struct {
	CategoryNumber  *int    `json:"category_number"`
	Name            *string `json:"name"`
	Producer        *string `json:"producer,omitempty"`
	Characteristics *string `json:"characteristics"`
}

// StoreProduct Models
type CreateStoreProductRequest struct {
	UPC                string  `json:"upc" binding:"required"`
	UPCProm            *string `json:"upc_prom,omitempty"`
	ProductID          int     `json:"product_id" binding:"required"`
	SellingPrice       float64 `json:"selling_price" binding:"required"`
	ProductsNumber     int     `json:"products_number" binding:"required"`
	PromotionalProduct bool    `json:"promotional_product"`
}

type UpdateStoreProductRequest struct {
	UPCProm            *string  `json:"upc_prom,omitempty"`
	ProductID          *int     `json:"product_id"`
	SellingPrice       *float64 `json:"selling_price"`
	ProductsNumber     *int     `json:"products_number"`
	PromotionalProduct *bool    `json:"promotional_product"`
}

// CustomerCard Models
type CreateCustomerCardRequest struct {
	CardNumber  string  `json:"card_number" binding:"required"`
	Surname     string  `json:"surname" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Patronymic  *string `json:"patronymic,omitempty"`
	PhoneNumber string  `json:"phone_number" binding:"required"`
	City        *string `json:"city,omitempty"`
	Street      *string `json:"street,omitempty"`
	ZipCode     *string `json:"zip_code,omitempty"`
	Percent     int     `json:"percent" binding:"required"`
}

type UpdateCustomerCardRequest struct {
	Surname     *string `json:"surname"`
	Name        *string `json:"name"`
	Patronymic  *string `json:"patronymic,omitempty"`
	PhoneNumber *string `json:"phone_number"`
	City        *string `json:"city,omitempty"`
	Street      *string `json:"street,omitempty"`
	ZipCode     *string `json:"zip_code,omitempty"`
	Percent     *int    `json:"percent"`
}

// Check Models
type CreateCheckRequest struct {
	Number     string    `json:"number" binding:"required"`
	EmployeeID string    `json:"employee_id" binding:"required"`
	CardNumber *string   `json:"card_number,omitempty"`
	PrintDate  time.Time `json:"print_date" binding:"required"`
	SumTotal   float64   `json:"sum_total" binding:"required"`
	VAT        float64   `json:"vat" binding:"required"`
}

// Sale Models
type CreateSaleRequest struct {
	UPC           string  `json:"upc" binding:"required"`
	CheckNumber   string  `json:"check_number" binding:"required"`
	ProductNumber int     `json:"product_number" binding:"required"`
	SellingPrice  float64 `json:"selling_price" binding:"required"`
}

// Auth Models
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
