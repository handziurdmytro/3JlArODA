package customercard

import (
	"context"
	"errors"
	"fmt"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	customercarddb "github.com/handziurdmytro/3JlArODA/business-service/internal/customercard/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	queries *customercarddb.Queries
}

func NewCustomerCardRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{queries: customercarddb.New(pool)}
}

func (r *Repository) Create(ctx context.Context, req CreateRequest) (*CustomerCard, error) {
	row, err := r.queries.CreateCustomerCard(ctx, customercarddb.CreateCustomerCardParams{
		CardNumber:     req.CardNumber,
		CustSurname:    req.Surname,
		CustName:       req.Name,
		CustPatronymic: textFromPtr(req.Patronymic),
		PhoneNumber:    req.PhoneNumber,
		City:           textFromPtr(req.City),
		Street:         textFromPtr(req.Street),
		ZipCode:        textFromPtr(req.ZipCode),
		Percent:        int32(req.Percent),
	})
	if err != nil {
		return nil, handleError(fmt.Sprintf("create customer card %q", req.CardNumber), err)
	}

	return mapCustomerCard(row), nil
}

func (r *Repository) Update(ctx context.Context, card CustomerCard) (*CustomerCard, error) {
	row, err := r.queries.UpdateCustomerCard(ctx, customercarddb.UpdateCustomerCardParams{
		CardNumber:     card.CardNumber,
		CustSurname:    card.Surname,
		CustName:       card.Name,
		CustPatronymic: textFromPtr(card.Patronymic),
		PhoneNumber:    card.PhoneNumber,
		City:           textFromPtr(card.City),
		Street:         textFromPtr(card.Street),
		ZipCode:        textFromPtr(card.ZipCode),
		Percent:        int32(card.Percent),
	})
	if err != nil {
		return nil, handleError(fmt.Sprintf("update customer card %q", card.CardNumber), err)
	}

	return mapCustomerCard(row), nil
}

func (r *Repository) Delete(ctx context.Context, cardNumber string) error {
	if _, err := r.GetByNumber(ctx, cardNumber); err != nil {
		return err
	}

	if err := r.queries.DeleteCustomerCard(ctx, cardNumber); err != nil {
		return handleError(fmt.Sprintf("delete customer card %q", cardNumber), err)
	}

	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]CustomerCard, error) {
	rows, err := r.queries.GetAllCustomerCards(ctx)
	if err != nil {
		return nil, handleError("get all customer cards", err)
	}

	return mapCustomerCards(rows), nil
}

func (r *Repository) GetByNumber(ctx context.Context, cardNumber string) (*CustomerCard, error) {
	row, err := r.queries.GetCustomerCardByNumber(ctx, cardNumber)
	if err != nil {
		return nil, handleError(fmt.Sprintf("get customer card by number %q", cardNumber), err)
	}

	return mapCustomerCard(row), nil
}

func (r *Repository) GetByPercent(ctx context.Context, percent int) ([]CustomerCard, error) {
	rows, err := r.queries.GetCustomerCardsByPercent(ctx, int32(percent))
	if err != nil {
		return nil, handleError(fmt.Sprintf("get customer cards by percent %d", percent), err)
	}

	return mapCustomerCards(rows), nil
}

func (r *Repository) SearchBySurname(ctx context.Context, surname string) ([]CustomerCard, error) {
	rows, err := r.queries.SearchCustomerCardsBySurname(ctx, textFromString(surname))
	if err != nil {
		return nil, handleError(fmt.Sprintf("search customer cards by surname %q", surname), err)
	}

	return mapCustomerCards(rows), nil
}

func mapCustomerCards(rows []customercarddb.CustomerCard) []CustomerCard {
	cards := make([]CustomerCard, 0, len(rows))
	for _, row := range rows {
		cards = append(cards, *mapCustomerCard(row))
	}

	return cards
}

func mapCustomerCard(row customercarddb.CustomerCard) *CustomerCard {
	return &CustomerCard{
		CardNumber:  row.CardNumber,
		Surname:     row.CustSurname,
		Name:        row.CustName,
		Patronymic:  ptrFromText(row.CustPatronymic),
		PhoneNumber: row.PhoneNumber,
		City:        ptrFromText(row.City),
		Street:      ptrFromText(row.Street),
		ZipCode:     ptrFromText(row.ZipCode),
		Percent:     int(row.Percent),
	}
}

func textFromString(value string) pgtype.Text {
	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}

func textFromPtr(value *string) pgtype.Text {
	if value == nil {
		return pgtype.Text{}
	}

	return textFromString(*value)
}

func ptrFromText(value pgtype.Text) *string {
	if !value.Valid {
		return nil
	}

	return &value.String
}

func handleError(operation string, err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("%s: %w", operation, common.ErrNotFound)
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return fmt.Errorf("%s: %w", operation, common.ErrAlreadyExists)
		case "23503":
			return fmt.Errorf("%s: %w", operation, common.ErrForeignKeyViolation)
		case "23514":
			return fmt.Errorf("%s: %w", operation, common.ErrCheckViolation)
		}
	}

	return fmt.Errorf("%s: %w", operation, err)
}
