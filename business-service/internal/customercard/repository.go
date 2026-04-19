package customercard

import (
	"context"
	"fmt"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	customercarddb "github.com/handziurdmytro/3JlArODA/business-service/internal/customercard/sqlc"
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
		CustPatronymic: common.TextFromPtr(req.Patronymic),
		PhoneNumber:    req.PhoneNumber,
		City:           common.TextFromPtr(req.City),
		Street:         common.TextFromPtr(req.Street),
		ZipCode:        common.TextFromPtr(req.ZipCode),
		Percent:        int32(req.Percent),
	})
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("create customer card %q", req.CardNumber), err)
	}

	return mapCustomerCard(row), nil
}

func (r *Repository) Update(ctx context.Context, card CustomerCard) (*CustomerCard, error) {
	row, err := r.queries.UpdateCustomerCard(ctx, customercarddb.UpdateCustomerCardParams{
		CardNumber:     card.CardNumber,
		CustSurname:    card.Surname,
		CustName:       card.Name,
		CustPatronymic: common.TextFromPtr(card.Patronymic),
		PhoneNumber:    card.PhoneNumber,
		City:           common.TextFromPtr(card.City),
		Street:         common.TextFromPtr(card.Street),
		ZipCode:        common.TextFromPtr(card.ZipCode),
		Percent:        int32(card.Percent),
	})
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("update customer card %q", card.CardNumber), err)
	}

	return mapCustomerCard(row), nil
}

func (r *Repository) Delete(ctx context.Context, cardNumber string) error {
	if _, err := r.GetByNumber(ctx, cardNumber); err != nil {
		return err
	}

	if err := r.queries.DeleteCustomerCard(ctx, cardNumber); err != nil {
		return common.MapRepositoryError(fmt.Sprintf("delete customer card %q", cardNumber), err)
	}

	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]CustomerCard, error) {
	rows, err := r.queries.GetAllCustomerCards(ctx)
	if err != nil {
		return nil, common.MapRepositoryError("get all customer cards", err)
	}

	return mapCustomerCards(rows), nil
}

func (r *Repository) GetByNumber(ctx context.Context, cardNumber string) (*CustomerCard, error) {
	row, err := r.queries.GetCustomerCardByNumber(ctx, cardNumber)
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get customer card by number %q", cardNumber), err)
	}

	return mapCustomerCard(row), nil
}

func (r *Repository) GetByPercent(ctx context.Context, percent int) ([]CustomerCard, error) {
	rows, err := r.queries.GetCustomerCardsByPercent(ctx, int32(percent))
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("get customer cards by percent %d", percent), err)
	}

	return mapCustomerCards(rows), nil
}

func (r *Repository) SearchBySurname(ctx context.Context, surname string) ([]CustomerCard, error) {
	rows, err := r.queries.SearchCustomerCardsBySurname(ctx, common.TextFromString(surname))
	if err != nil {
		return nil, common.MapRepositoryError(fmt.Sprintf("search customer cards by surname %q", surname), err)
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
		Patronymic:  common.PtrFromText(row.CustPatronymic),
		PhoneNumber: row.PhoneNumber,
		City:        common.PtrFromText(row.City),
		Street:      common.PtrFromText(row.Street),
		ZipCode:     common.PtrFromText(row.ZipCode),
		Percent:     int(row.Percent),
	}
}
