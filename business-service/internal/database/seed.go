package database

import (
	"context"
	"log/slog"
	"time"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/employee"
)

func Seed(ctx context.Context, repo *employee.Repository) {
	adminID := "ADM-000000"

	_, err := repo.GetEmployeeByID(ctx, adminID)
	if err == nil {
		slog.Info("Default admin profile already exists in business DB, skipping seed")
		return
	}

	patronymic := "Adminovych"
	adminReq := employee.CreateRequest{
		ID:          adminID,
		Surname:     "Adminenko",
		Name:        "Admin",
		Patronymic:  &patronymic,
		Role:        "manager",
		Salary:      100000.0,
		DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		DateOfStart: time.Now(),
		PhoneNumber: "+380000000000",
		City:        "Kyiv",
		Street:      "Volodymyrska",
		ZipCode:     "01001",
	}

	err = repo.CreateEmployee(ctx, adminReq)
	if err != nil {
		slog.Error("CRITICAL: Failed to seed default admin in business DB", "error", err)
		return
	}

	slog.Info("Successfully seeded default manager profile", "id", adminID)
}
