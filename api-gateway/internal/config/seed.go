package config

import (
	"context"
	"log"
	"time"

	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/grpc/clients"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/handlers"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	authpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb"
)

func Execute(authClient *clients.Auth, empClient handlers.EmployeeClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	adminID := "ADM-000000"
	_, err := empClient.GetByID(ctx, adminID)
	if err == nil {
		log.Println("[INFO] Default admin already exists. Skipping.")
	} else {
		log.Println("[INFO] Seeding default admin entity...")
		_, err = authClient.Register(ctx, &authpb.RegisterRequest{
			Id:       adminID,
			Username: "zlahoda@ukma.edu.ua",
			Password: "secret secret",
			Role:     "manager",
		})
		if err != nil {
			log.Printf("[ERROR] Failed to seed admin in auth service: %v", err)
		} else {
			patronymic := "Adminovych"
			_, err = empClient.Create(ctx, models.CreateEmployeeRequest{
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
			})
			if err != nil {
				log.Printf("[ERROR] Failed to seed admin in business service: %v", err)
			}
		}
	}

	cashierID := "CSH-000000"
	_, err = empClient.GetByID(ctx, cashierID)
	if err == nil {
		log.Println("[INFO] Default cashier already exists. Skipping.")
		return
	}

	log.Println("[INFO] Seeding default cashier entity...")
	_, err = authClient.Register(ctx, &authpb.RegisterRequest{
		Id:       cashierID,
		Username: "cashier@ukma.edu.ua",
		Password: "cash cash",
		Role:     "cashier",
	})
	if err != nil {
		log.Printf("[ERROR] Failed to seed cashier in auth service: %v", err)
		return
	}

	patronymic2 := "Kasyrovych"
	_, err = empClient.Create(ctx, models.CreateEmployeeRequest{
		ID:          cashierID,
		Surname:     "Kasyrenko",
		Name:        "Kasyr",
		Patronymic:  &patronymic2,
		Role:        "cashier",
		Salary:      50000.0,
		DateOfBirth: time.Date(2001, 3, 2, 0, 0, 0, 0, time.UTC),
		DateOfStart: time.Now(),
		PhoneNumber: "+380999999999",
		City:        "Lviv",
		Street:      "Kulparkivska",
		ZipCode:     "28190",
	})
	if err != nil {
		log.Printf("[ERROR] Failed to seed cashier in business service: %v", err)
		return
	}

	log.Println("[INFO] Successfully split and seeded default accounts!")
}
