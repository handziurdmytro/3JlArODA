package database

import (
	"context"
	"errors"
	"log/slog"
	"math"
	"time"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/category"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/check"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/customercard"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/employee"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/product"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/sale"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/storeproduct"
)

type demoSaleLine struct {
	upc      string
	quantity int
	price    float64
}

func Seed(
	ctx context.Context,
	employeeRepo *employee.Repository,
	categoryRepo *category.Repository,
	productRepo *product.Repository,
	storeProductRepo *storeproduct.Repository,
	customerCardRepo *customercard.Repository,
	checkRepo *check.Repository,
	saleRepo *sale.Repository,
) {
	seedDefaultAdmin(ctx, employeeRepo)
	seedDemoData(ctx, employeeRepo, categoryRepo, productRepo, storeProductRepo, customerCardRepo, checkRepo, saleRepo)
}

func seedDefaultAdmin(ctx context.Context, repo *employee.Repository) {
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

func seedDemoData(
	ctx context.Context,
	employeeRepo *employee.Repository,
	categoryRepo *category.Repository,
	productRepo *product.Repository,
	storeProductRepo *storeproduct.Repository,
	customerCardRepo *customercard.Repository,
	checkRepo *check.Repository,
	saleRepo *sale.Repository,
) {
	if err := seedDemoEmployees(ctx, employeeRepo); err != nil {
		slog.Error("failed to seed demo employees", "error", err)
		return
	}

	categories, err := seedDemoCategories(ctx, categoryRepo)
	if err != nil {
		slog.Error("failed to seed demo categories", "error", err)
		return
	}

	products, err := seedDemoProducts(ctx, productRepo, categories)
	if err != nil {
		slog.Error("failed to seed demo products", "error", err)
		return
	}

	if err := seedDemoStoreProducts(ctx, storeProductRepo, products); err != nil {
		slog.Error("failed to seed demo store products", "error", err)
		return
	}

	if err := seedDemoCustomerCards(ctx, customerCardRepo); err != nil {
		slog.Error("failed to seed demo customer cards", "error", err)
		return
	}

	if err := seedDemoChecksAndSales(ctx, checkRepo, saleRepo); err != nil {
		slog.Error("failed to seed demo checks and sales", "error", err)
		return
	}

	slog.Info("Successfully seeded demo business data")
}

func seedDemoEmployees(ctx context.Context, repo *employee.Repository) error {
	patronymicCashier := "Mykolayovych"
	patronymicManager := "Oleksandrivna"
	employees := []employee.CreateRequest{
		{
			ID:          "CSH-000001",
			Surname:     "Kovalchuk",
			Name:        "Mykola",
			Patronymic:  &patronymicCashier,
			Role:        "cashier",
			Salary:      42000,
			DateOfBirth: time.Date(1997, 5, 14, 0, 0, 0, 0, time.UTC),
			DateOfStart: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
			PhoneNumber: "+380671112201",
			City:        "Kyiv",
			Street:      "Sichovykh Striltsiv 12",
			ZipCode:     "04053",
		},
		{
			ID:          "CSH-000002",
			Surname:     "Shevchenko",
			Name:        "Iryna",
			Patronymic:  nil,
			Role:        "cashier",
			Salary:      43500,
			DateOfBirth: time.Date(1999, 9, 3, 0, 0, 0, 0, time.UTC),
			DateOfStart: time.Date(2024, 6, 10, 0, 0, 0, 0, time.UTC),
			PhoneNumber: "+380671112202",
			City:        "Lviv",
			Street:      "Horodotska 45",
			ZipCode:     "79007",
		},
		{
			ID:          "MGR-000001",
			Surname:     "Melnyk",
			Name:        "Oksana",
			Patronymic:  &patronymicManager,
			Role:        "manager",
			Salary:      75000,
			DateOfBirth: time.Date(1992, 11, 20, 0, 0, 0, 0, time.UTC),
			DateOfStart: time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
			PhoneNumber: "+380671112203",
			City:        "Kyiv",
			Street:      "Yaroslaviv Val 8",
			ZipCode:     "01054",
		},
	}

	for _, req := range employees {
		if _, err := repo.GetEmployeeByID(ctx, req.ID); err == nil {
			continue
		}
		if err := repo.CreateEmployee(ctx, req); err != nil && !errors.Is(err, common.ErrAlreadyExists) {
			return err
		}
	}

	return nil
}

func seedDemoCategories(ctx context.Context, repo *category.Repository) (map[string]int, error) {
	categoryNames := []string{"Bakery", "Dairy", "Drinks", "Household", "Meat"}
	result := make(map[string]int, len(categoryNames))

	for _, name := range categoryNames {
		existing, err := repo.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		for _, item := range existing {
			if item.Name == name {
				result[name] = item.Number
				break
			}
		}
		if result[name] != 0 {
			continue
		}

		created, err := repo.Create(ctx, category.CreateRequest{Name: name})
		if err != nil && !errors.Is(err, common.ErrAlreadyExists) {
			return nil, err
		}
		if err == nil {
			result[name] = created.Number
		}
	}

	return result, nil
}

func seedDemoProducts(ctx context.Context, repo *product.Repository, categories map[string]int) (map[string]int, error) {
	type demoProduct struct {
		key             string
		categoryName    string
		name            string
		producer        string
		characteristics string
	}

	products := []demoProduct{
		{key: "milk", categoryName: "Dairy", name: "Milk 2.5%", producer: "Halychyna", characteristics: "1 liter bottle"},
		{key: "yogurt", categoryName: "Dairy", name: "Greek Yogurt", producer: "Molokiya", characteristics: "Plain yogurt, 300 g"},
		{key: "bread", categoryName: "Bakery", name: "Rye Bread", producer: "Kyivkhlib", characteristics: "Dark rye bread, 500 g"},
		{key: "chicken", categoryName: "Meat", name: "Chicken Fillet", producer: "Nasha Riaba", characteristics: "Chilled fillet, 1 kg"},
		{key: "water", categoryName: "Drinks", name: "Mineral Water", producer: "Morshynska", characteristics: "Still water, 1.5 liter bottle"},
		{key: "soap", categoryName: "Household", name: "Dish Soap", producer: "Helper", characteristics: "Lemon scent, 500 ml"},
	}

	result := make(map[string]int, len(products))
	existing, err := repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, req := range products {
		for _, item := range existing {
			if item.Name == req.name {
				result[req.key] = item.ID
				break
			}
		}
		if result[req.key] != 0 {
			continue
		}

		categoryNumber := categories[req.categoryName]
		if categoryNumber == 0 {
			return nil, common.ErrNotFound
		}

		producer := req.producer
		created, err := repo.Create(ctx, product.CreateRequest{
			CategoryNumber:  categoryNumber,
			Name:            req.name,
			Producer:        &producer,
			Characteristics: req.characteristics,
		})
		if err != nil && !errors.Is(err, common.ErrAlreadyExists) {
			return nil, err
		}
		if err == nil {
			result[req.key] = created.ID
		}
	}

	return result, nil
}

func seedDemoStoreProducts(ctx context.Context, repo *storeproduct.Repository, products map[string]int) error {
	milkRegular := "482000000001"
	breadRegular := "482000000003"
	items := []storeproduct.CreateRequest{
		{UPC: milkRegular, ProductID: products["milk"], SellingPrice: 42.50, ProductsNumber: 120, PromotionalProduct: false},
		{UPC: "482000000002", ProductID: products["yogurt"], SellingPrice: 38.90, ProductsNumber: 80, PromotionalProduct: false},
		{UPC: breadRegular, ProductID: products["bread"], SellingPrice: 31.00, ProductsNumber: 60, PromotionalProduct: false},
		{UPC: "482000000004", ProductID: products["chicken"], SellingPrice: 189.90, ProductsNumber: 35, PromotionalProduct: false},
		{UPC: "482000000005", ProductID: products["water"], SellingPrice: 24.50, ProductsNumber: 200, PromotionalProduct: false},
		{UPC: "482000000006", ProductID: products["soap"], SellingPrice: 56.00, ProductsNumber: 45, PromotionalProduct: false},
		{UPC: "482000000101", UPCProm: &milkRegular, ProductID: products["milk"], SellingPrice: 36.90, ProductsNumber: 40, PromotionalProduct: true},
		{UPC: "482000000103", UPCProm: &breadRegular, ProductID: products["bread"], SellingPrice: 25.90, ProductsNumber: 25, PromotionalProduct: true},
	}

	for _, req := range items {
		if req.ProductID == 0 {
			return common.ErrNotFound
		}
		if _, err := repo.GetByUPC(ctx, req.UPC); err == nil {
			continue
		}
		if _, err := repo.Create(ctx, req); err != nil && !errors.Is(err, common.ErrAlreadyExists) {
			return err
		}
	}

	return nil
}

func seedDemoCustomerCards(ctx context.Context, repo *customercard.Repository) error {
	city := "Kyiv"
	street := "Khreshchatyk 22"
	zip := "01001"
	patronymic := "Petrivna"
	cards := []customercard.CreateRequest{
		{CardNumber: "CARD000000001", Surname: "Bondarenko", Name: "Tetiana", Patronymic: &patronymic, PhoneNumber: "+380501110001", City: &city, Street: &street, ZipCode: &zip, Percent: 3},
		{CardNumber: "CARD000000002", Surname: "Kravchenko", Name: "Andrii", PhoneNumber: "+380501110002", City: &city, Street: &street, ZipCode: &zip, Percent: 5},
		{CardNumber: "CARD000000003", Surname: "Tkachenko", Name: "Marta", PhoneNumber: "+380501110003", City: &city, Street: &street, ZipCode: &zip, Percent: 10},
		{CardNumber: "CARD000000004", Surname: "Lysenko", Name: "Serhii", PhoneNumber: "+380501110004", City: &city, Street: &street, ZipCode: &zip, Percent: 0},
		{CardNumber: "CARD000000005", Surname: "Polishchuk", Name: "Nadiia", PhoneNumber: "+380501110005", City: &city, Street: &street, ZipCode: &zip, Percent: 15},
	}

	for _, req := range cards {
		if _, err := repo.GetByNumber(ctx, req.CardNumber); err == nil {
			continue
		}
		if _, err := repo.Create(ctx, req); err != nil && !errors.Is(err, common.ErrAlreadyExists) {
			return err
		}
	}

	return nil
}

func seedDemoChecksAndSales(ctx context.Context, checkRepo *check.Repository, saleRepo *sale.Repository) error {
	type demoCheck struct {
		number     string
		employeeID string
		cardNumber *string
		printDate  time.Time
		lines      []demoSaleLine
	}

	cardOne := "CARD000000001"
	cardTwo := "CARD000000002"
	cardThree := "CARD000000003"
	cardFive := "CARD000000005"
	checks := []demoCheck{
		{
			number:     "CHK0000001",
			employeeID: "CSH-000001",
			cardNumber: &cardOne,
			printDate:  time.Date(2026, 4, 18, 10, 20, 0, 0, time.UTC),
			lines: []demoSaleLine{
				{upc: "482000000001", quantity: 2, price: 42.50},
				{upc: "482000000003", quantity: 1, price: 31.00},
			},
		},
		{
			number:     "CHK0000002",
			employeeID: "CSH-000001",
			cardNumber: &cardTwo,
			printDate:  time.Date(2026, 4, 18, 14, 5, 0, 0, time.UTC),
			lines: []demoSaleLine{
				{upc: "482000000101", quantity: 3, price: 36.90},
				{upc: "482000000005", quantity: 2, price: 24.50},
			},
		},
		{
			number:     "CHK0000003",
			employeeID: "CSH-000002",
			cardNumber: nil,
			printDate:  time.Date(2026, 4, 19, 9, 45, 0, 0, time.UTC),
			lines: []demoSaleLine{
				{upc: "482000000004", quantity: 1, price: 189.90},
				{upc: "482000000002", quantity: 2, price: 38.90},
			},
		},
		{
			number:     "CHK0000004",
			employeeID: "CSH-000002",
			cardNumber: &cardThree,
			printDate:  time.Date(2026, 4, 20, 16, 30, 0, 0, time.UTC),
			lines: []demoSaleLine{
				{upc: "482000000006", quantity: 1, price: 56.00},
				{upc: "482000000103", quantity: 2, price: 25.90},
			},
		},
		{
			number:     "CHK0000005",
			employeeID: "CSH-000001",
			cardNumber: &cardFive,
			printDate:  time.Date(2026, 4, 21, 11, 15, 0, 0, time.UTC),
			lines: []demoSaleLine{
				{upc: "482000000001", quantity: 1, price: 42.50},
				{upc: "482000000002", quantity: 1, price: 38.90},
				{upc: "482000000005", quantity: 4, price: 24.50},
			},
		},
	}

	for _, req := range checks {
		sumTotal := checkTotal(req.lines)
		if _, err := checkRepo.GetByNumber(ctx, req.number); err != nil {
			if err := checkRepo.Create(ctx, check.CreateRequest{
				Number:     req.number,
				EmployeeID: req.employeeID,
				CardNumber: req.cardNumber,
				PrintDate:  req.printDate,
				SumTotal:   sumTotal,
				VAT:        roundMoney(sumTotal * 0.2),
			}); err != nil && !errors.Is(err, common.ErrAlreadyExists) {
				return err
			}
		}

		for _, item := range req.lines {
			err := saleRepo.Create(ctx, sale.CreateRequest{
				UPC:           item.upc,
				CheckNumber:   req.number,
				ProductNumber: item.quantity,
				SellingPrice:  item.price,
			})
			if err != nil && !errors.Is(err, common.ErrAlreadyExists) {
				return err
			}
		}
	}

	return nil
}

func checkTotal(lines []demoSaleLine) float64 {
	var total float64
	for _, line := range lines {
		total += float64(line.quantity) * line.price
	}

	return roundMoney(total)
}

func roundMoney(value float64) float64 {
	return math.Round(value*100) / 100
}
