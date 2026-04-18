package clients

import (
	"context"

	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	customercardpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/customercard"
	"google.golang.org/grpc"
)

type CustomerCardClient struct {
	client customercardpb.CustomerCardServiceClient
}

func NewCustomerCardClient(conn grpc.ClientConnInterface) *CustomerCardClient {
	return &CustomerCardClient{
		client: customercardpb.NewCustomerCardServiceClient(conn),
	}
}

func (cl *CustomerCardClient) Create(ctx context.Context, req models.CreateCustomerCardRequest) (*customercardpb.CustomerCard, error) {
	resp, err := cl.client.CreateCustomerCard(ctx, &customercardpb.CreateCustomerCardRequest{
		CardNumber:  req.CardNumber,
		Surname:     req.Surname,
		Name:        req.Name,
		Patronymic:  req.Patronymic,
		PhoneNumber: req.PhoneNumber,
		City:        req.City,
		Street:      req.Street,
		ZipCode:     req.ZipCode,
		Percent:     int32(req.Percent),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetCustomerCard(), nil
}

func (cl *CustomerCardClient) GetByNumber(ctx context.Context, cardNumber string) (*customercardpb.CustomerCard, error) {
	resp, err := cl.client.GetCustomerCard(ctx, &customercardpb.GetCustomerCardRequest{
		CardNumber: cardNumber,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetCustomerCard(), nil
}

func (cl *CustomerCardClient) GetAll(ctx context.Context) ([]*customercardpb.CustomerCard, error) {
	resp, err := cl.client.GetAllCustomerCards(ctx, &customercardpb.GetAllCustomerCardsRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetCustomerCards(), nil
}

func (cl *CustomerCardClient) GetByPercent(ctx context.Context, percent int) ([]*customercardpb.CustomerCard, error) {
	resp, err := cl.client.GetCustomerCardsByPercent(ctx, &customercardpb.GetCustomerCardsByPercentRequest{
		Percent: int32(percent),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetCustomerCards(), nil
}

func (cl *CustomerCardClient) SearchBySurname(ctx context.Context, surname string) ([]*customercardpb.CustomerCard, error) {
	resp, err := cl.client.SearchCustomerCardsBySurname(ctx, &customercardpb.SearchCustomerCardsBySurnameRequest{
		Surname: surname,
	})
	if err != nil {
		return nil, err
	}

	return resp.GetCustomerCards(), nil
}

func (cl *CustomerCardClient) Update(ctx context.Context, cardNumber string, req models.UpdateCustomerCardRequest) (*customercardpb.CustomerCard, error) {
	resp, err := cl.client.UpdateCustomerCard(ctx, &customercardpb.UpdateCustomerCardRequest{
		CardNumber:  cardNumber,
		Surname:     req.Surname,
		Name:        req.Name,
		Patronymic:  req.Patronymic,
		PhoneNumber: req.PhoneNumber,
		City:        req.City,
		Street:      req.Street,
		ZipCode:     req.ZipCode,
		Percent:     int32(req.Percent),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetCustomerCard(), nil
}

func (cl *CustomerCardClient) Delete(ctx context.Context, cardNumber string) error {
	_, err := cl.client.DeleteCustomerCard(ctx, &customercardpb.DeleteCustomerCardRequest{
		CardNumber: cardNumber,
	})

	return err
}
