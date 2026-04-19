package check

import (
	"context"
	"fmt"
	"time"

	checkdb "github.com/handziurdmytro/3JlArODA/business-service/internal/check/sqlc"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	queries *checkdb.Queries
}

func NewCheckRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{queries: checkdb.New(pool)}
}

func (r *Repository) Create(ctx context.Context, req CreateRequest) error {
	sumTotal, err := common.NumericFromFloat64(req.SumTotal)
	if err != nil {
		return fmt.Errorf("create check %q: %w", req.Number, err)
	}

	vat, err := common.NumericFromFloat64(req.VAT)
	if err != nil {
		return fmt.Errorf("create check %q: %w", req.Number, err)
	}

	err = r.queries.CreateCheck(ctx, checkdb.CreateCheckParams{
		CheckNumber: req.Number,
		IDEmployee:  req.EmployeeID,
		CardNumber:  common.TextFromPtr(req.CardNumber),
		PrintDate:   common.TimestampFromTime(req.PrintDate),
		SumTotal:    sumTotal,
		Vat:         vat,
	})
	if err != nil {
		return common.MapRepositoryError(fmt.Sprintf("create check %q", req.Number), err)
	}

	return nil
}

func (r *Repository) DeleteByNumber(ctx context.Context, number string) error {
	if _, err := r.GetByNumber(ctx, number); err != nil {
		return err
	}

	if err := r.queries.DeleteCheckByNumber(ctx, number); err != nil {
		return common.MapRepositoryError(fmt.Sprintf("delete check %q", number), err)
	}

	return nil
}

func (r *Repository) GetByNumber(ctx context.Context, number string) (*Check, error) {
	row, err := r.queries.GetCheckByNumber(ctx, number)
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get check by number %q", number), err)
	}

	return mapCheck(row)
}

func (r *Repository) GetAll(ctx context.Context) ([]Check, error) {
	rows, err := r.queries.GetAllChecks(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get all checks", err)
	}

	return mapChecks(rows)
}

func (r *Repository) GetAllOfTheDayByCashier(ctx context.Context, employeeID string, day time.Time) ([]Check, error) {
	rows, err := r.queries.GetAllChecksOfTheDayByCashier(ctx, checkdb.GetAllChecksOfTheDayByCashierParams{
		IDEmployee: employeeID,
		PrintDate:  common.TimestampFromTime(day),
	})
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get checks of day by cashier %q", employeeID), err)
	}

	return mapChecks(rows)
}

func (r *Repository) GetAllOfThePeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) ([]Check, error) {
	rows, err := r.queries.GetAllChecksOfThePeriodByCashier(ctx, checkdb.GetAllChecksOfThePeriodByCashierParams{
		IDEmployee:  employeeID,
		PrintDate:   common.TimestampFromTime(from),
		PrintDate_2: common.TimestampFromTime(to),
	})
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get checks of period by cashier %q", employeeID), err)
	}

	return mapChecks(rows)
}

func (r *Repository) GetFullDataByNumber(ctx context.Context, number string) ([]FullCheckData, error) {
	rows, err := r.queries.GetFullCheckDataByID(ctx, number)
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get full check data by number %q", number), err)
	}
	if len(rows) == 0 {
		return nil, common.MapRepositoryError(fmt.Sprintf("get full check data by number %q", number), common.ErrNotFound)
	}

	return mapFullCheckDataRows(rows)
}

func (r *Repository) GetDetailsOfThePeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) ([]Detail, error) {
	rows, err := r.queries.GetCheckDetailsOfThePeriodByCashier(ctx, checkdb.GetCheckDetailsOfThePeriodByCashierParams{
		IDEmployee:  employeeID,
		PrintDate:   common.TimestampFromTime(from),
		PrintDate_2: common.TimestampFromTime(to),
	})
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get check details of period by cashier %q", employeeID), err)
	}

	return mapDetailsByCashier(rows)
}

func (r *Repository) GetDetailsOfThePeriod(ctx context.Context, from, to time.Time) ([]Detail, error) {
	rows, err := r.queries.GetCheckDetailsOfThePeriod(ctx, checkdb.GetCheckDetailsOfThePeriodParams{
		PrintDate:   common.TimestampFromTime(from),
		PrintDate_2: common.TimestampFromTime(to),
	})
	if err != nil {
		return nil, common.MapRepositoryError("get check details of period", err)
	}

	return mapDetails(rows)
}

func (r *Repository) GetSumOfThePeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) (float64, error) {
	total, err := r.queries.GetSumOfAllChecksOfThePeriodByCashier(ctx, checkdb.GetSumOfAllChecksOfThePeriodByCashierParams{
		IDEmployee:  employeeID,
		PrintDate:   common.TimestampFromTime(from),
		PrintDate_2: common.TimestampFromTime(to),
	})
	if err != nil {
		return 0, common.MapRepositoryError(fmt.Sprintf("get check sum of period by cashier %q", employeeID), err)
	}

	return common.Float64FromNumeric(total)
}

func (r *Repository) GetSumOfThePeriod(ctx context.Context, from, to time.Time) (float64, error) {
	total, err := r.queries.GetSumOfAllChecksOfThePeriod(ctx, checkdb.GetSumOfAllChecksOfThePeriodParams{
		PrintDate:   common.TimestampFromTime(from),
		PrintDate_2: common.TimestampFromTime(to),
	})
	if err != nil {
		return 0, common.MapRepositoryError("get check sum of period", err)
	}

	return common.Float64FromNumeric(total)
}

func mapChecks(rows []checkdb.Check) ([]Check, error) {
	checks := make([]Check, 0, len(rows))
	for _, row := range rows {
		check, err := mapCheck(row)
		if err != nil {
			return nil, err
		}
		checks = append(checks, *check)
	}

	return checks, nil
}

func mapCheck(row checkdb.Check) (*Check, error) {
	sumTotal, err := common.Float64FromNumeric(row.SumTotal)
	if err != nil {
		return nil, fmt.Errorf("map check sum total: %w", err)
	}

	vat, err := common.Float64FromNumeric(row.Vat)
	if err != nil {
		return nil, fmt.Errorf("map check vat: %w", err)
	}

	return &Check{
		Number:     row.CheckNumber,
		EmployeeID: row.IDEmployee,
		CardNumber: common.PtrFromText(row.CardNumber),
		PrintDate:  common.TimeFromTimestamp(row.PrintDate),
		SumTotal:   sumTotal,
		VAT:        vat,
	}, nil
}

func mapFullCheckDataRows(rows []checkdb.GetFullCheckDataByIDRow) ([]FullCheckData, error) {
	checks := make([]FullCheckData, 0, len(rows))
	for _, row := range rows {
		sumTotal, err := common.Float64FromNumeric(row.SumTotal)
		if err != nil {
			return nil, fmt.Errorf("map full check sum total: %w", err)
		}

		vat, err := common.Float64FromNumeric(row.Vat)
		if err != nil {
			return nil, fmt.Errorf("map full check vat: %w", err)
		}

		sellingPrice, err := common.Float64FromNumeric(row.SellingPrice)
		if err != nil {
			return nil, fmt.Errorf("map full check selling price: %w", err)
		}

		checks = append(checks, FullCheckData{
			CheckNumber:        row.CheckNumber,
			PrintDate:          common.TimeFromTimestamp(row.PrintDate),
			SumTotal:           sumTotal,
			VAT:                vat,
			UPC:                row.Upc,
			Quantity:           int(row.Quantity),
			SellingPrice:       sellingPrice,
			ProductName:        row.ProductName,
			EmployeeSurname:    row.EmplSurname,
			EmployeeName:       row.EmplName,
			EmployeePatronymic: common.PtrFromText(row.EmplPatronymic),
		})
	}

	return checks, nil
}

func mapDetailsByCashier(rows []checkdb.GetCheckDetailsOfThePeriodByCashierRow) ([]Detail, error) {
	details := make([]Detail, 0, len(rows))
	for _, row := range rows {
		detail, err := mapDetail(row.CheckNumber, row.PrintDate, row.SumTotal, row.ProductName, row.Upc, row.Quantity, row.SellingPrice)
		if err != nil {
			return nil, err
		}
		details = append(details, *detail)
	}

	return details, nil
}

func mapDetails(rows []checkdb.GetCheckDetailsOfThePeriodRow) ([]Detail, error) {
	details := make([]Detail, 0, len(rows))
	for _, row := range rows {
		detail, err := mapDetail(row.CheckNumber, row.PrintDate, row.SumTotal, row.ProductName, row.Upc, row.Quantity, row.SellingPrice)
		if err != nil {
			return nil, err
		}
		details = append(details, *detail)
	}

	return details, nil
}

func mapDetail(
	checkNumber string,
	printDate pgtype.Timestamp,
	sumTotal pgtype.Numeric,
	productName string,
	upc string,
	quantity int32,
	sellingPriceValue pgtype.Numeric,
) (*Detail, error) {
	total, err := common.Float64FromNumeric(sumTotal)
	if err != nil {
		return nil, fmt.Errorf("map check detail sum total: %w", err)
	}

	sellingPrice, err := common.Float64FromNumeric(sellingPriceValue)
	if err != nil {
		return nil, fmt.Errorf("map check detail selling price: %w", err)
	}

	return &Detail{
		CheckNumber:  checkNumber,
		PrintDate:    common.TimeFromTimestamp(printDate),
		SumTotal:     total,
		ProductName:  productName,
		UPC:          upc,
		Quantity:     int(quantity),
		SellingPrice: sellingPrice,
	}, nil
}
