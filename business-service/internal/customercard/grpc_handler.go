package customercard

import (
	"context"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	customercardpb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/customercard"
)

type GRPCHandler struct {
	customercardpb.UnimplementedCustomerCardServiceServer

	service *Service
}

func NewGRPCHandler(service *Service) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (h *GRPCHandler) CreateCustomerCard(ctx context.Context, req *customercardpb.CreateCustomerCardRequest) (*customercardpb.CustomerCardResponse, error) {
	card, err := h.service.Create(ctx, CreateRequest{
		CardNumber:  req.GetCardNumber(),
		Surname:     req.GetSurname(),
		Name:        req.GetName(),
		Patronymic:  req.Patronymic,
		PhoneNumber: req.GetPhoneNumber(),
		City:        req.City,
		Street:      req.Street,
		ZipCode:     req.ZipCode,
		Percent:     int(req.GetPercent()),
	})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &customercardpb.CustomerCardResponse{CustomerCard: toProtoCustomerCard(card)}, nil
}

func (h *GRPCHandler) UpdateCustomerCard(ctx context.Context, req *customercardpb.UpdateCustomerCardRequest) (*customercardpb.CustomerCardResponse, error) {
	card, err := h.service.Update(ctx, CustomerCard{
		CardNumber:  req.GetCardNumber(),
		Surname:     req.GetSurname(),
		Name:        req.GetName(),
		Patronymic:  req.Patronymic,
		PhoneNumber: req.GetPhoneNumber(),
		City:        req.City,
		Street:      req.Street,
		ZipCode:     req.ZipCode,
		Percent:     int(req.GetPercent()),
	})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &customercardpb.CustomerCardResponse{CustomerCard: toProtoCustomerCard(card)}, nil
}

func (h *GRPCHandler) DeleteCustomerCard(ctx context.Context, req *customercardpb.DeleteCustomerCardRequest) (*customercardpb.DeleteCustomerCardResponse, error) {
	if err := h.service.Delete(ctx, req.GetCardNumber()); err != nil {
		return nil, common.ToStatusError(err)
	}

	return &customercardpb.DeleteCustomerCardResponse{Success: true}, nil
}

func (h *GRPCHandler) GetCustomerCard(ctx context.Context, req *customercardpb.GetCustomerCardRequest) (*customercardpb.CustomerCardResponse, error) {
	card, err := h.service.GetByNumber(ctx, req.GetCardNumber())
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &customercardpb.CustomerCardResponse{CustomerCard: toProtoCustomerCard(card)}, nil
}

func (h *GRPCHandler) GetAllCustomerCards(ctx context.Context, _ *customercardpb.GetAllCustomerCardsRequest) (*customercardpb.GetCustomerCardsResponse, error) {
	cards, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &customercardpb.GetCustomerCardsResponse{CustomerCards: toProtoCustomerCards(cards)}, nil
}

func (h *GRPCHandler) GetCustomerCardsByPercent(ctx context.Context, req *customercardpb.GetCustomerCardsByPercentRequest) (*customercardpb.GetCustomerCardsResponse, error) {
	cards, err := h.service.GetByPercent(ctx, int(req.GetPercent()))
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &customercardpb.GetCustomerCardsResponse{CustomerCards: toProtoCustomerCards(cards)}, nil
}

func (h *GRPCHandler) SearchCustomerCardsBySurname(ctx context.Context, req *customercardpb.SearchCustomerCardsBySurnameRequest) (*customercardpb.GetCustomerCardsResponse, error) {
	cards, err := h.service.SearchBySurname(ctx, req.GetSurname())
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &customercardpb.GetCustomerCardsResponse{CustomerCards: toProtoCustomerCards(cards)}, nil
}

func (h *GRPCHandler) GetCustomerCardsWhoBoughtAllProductsFromNonEmptyCategory(ctx context.Context, req *customercardpb.GetCustomerCardsWhoBoughtAllProductsFromNonEmptyCategoryRequest) (*customercardpb.GetCustomerCardsResponse, error) {
	from, err := common.ParseProtoTime(req.GetFrom())
	if err != nil {
		return nil, err
	}
	to, err := common.ParseProtoTime(req.GetTo())
	if err != nil {
		return nil, err
	}

	cards, err := h.service.GetWhoBoughtAllProductsFromCategory(ctx, int(req.GetCategoryNumber()), from, to)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &customercardpb.GetCustomerCardsResponse{CustomerCards: toProtoCustomerCards(cards)}, nil
}

func toProtoCustomerCards(cards []CustomerCard) []*customercardpb.CustomerCard {
	result := make([]*customercardpb.CustomerCard, 0, len(cards))
	for _, card := range cards {
		card := card
		result = append(result, toProtoCustomerCard(&card))
	}

	return result
}

func toProtoCustomerCard(card *CustomerCard) *customercardpb.CustomerCard {
	if card == nil {
		return nil
	}

	return &customercardpb.CustomerCard{
		CardNumber:  card.CardNumber,
		Surname:     card.Surname,
		Name:        card.Name,
		Patronymic:  card.Patronymic,
		PhoneNumber: card.PhoneNumber,
		City:        card.City,
		Street:      card.Street,
		ZipCode:     card.ZipCode,
		Percent:     int32(card.Percent),
	}
}
