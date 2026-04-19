package check

import (
	"context"
	"time"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/common"
	checkpb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/check"
)

type GRPCHandler struct {
	checkpb.UnimplementedCheckServiceServer

	service *Service
}

func NewGRPCHandler(service *Service) *GRPCHandler {
	return &GRPCHandler{service: service}
}

func (h *GRPCHandler) CreateCheck(ctx context.Context, req *checkpb.CreateCheckRequest) (*checkpb.CreateCheckResponse, error) {
	printDate, err := common.ParseProtoTime(req.GetPrintDate())
	if err != nil {
		return nil, err
	}

	err = h.service.Create(ctx, CreateRequest{
		Number:     req.GetNumber(),
		EmployeeID: req.GetEmployeeId(),
		CardNumber: req.CardNumber,
		PrintDate:  printDate,
		SumTotal:   req.GetSumTotal(),
		VAT:        req.GetVat(),
	})
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &checkpb.CreateCheckResponse{Success: true}, nil
}

func (h *GRPCHandler) DeleteCheck(ctx context.Context, req *checkpb.DeleteCheckRequest) (*checkpb.DeleteCheckResponse, error) {
	if err := h.service.DeleteByNumber(ctx, req.GetNumber()); err != nil {
		return nil, common.ToStatusError(err)
	}

	return &checkpb.DeleteCheckResponse{Success: true}, nil
}

func (h *GRPCHandler) GetAllChecks(ctx context.Context, _ *checkpb.GetAllChecksRequest) (*checkpb.GetChecksResponse, error) {
	checks, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &checkpb.GetChecksResponse{Checks: toProtoChecks(checks)}, nil
}

func (h *GRPCHandler) GetChecksOfTheDayByCashier(ctx context.Context, req *checkpb.GetChecksOfTheDayByCashierRequest) (*checkpb.GetChecksResponse, error) {
	day, err := common.ParseProtoTime(req.GetDate())
	if err != nil {
		return nil, err
	}

	checks, err := h.service.GetAllOfTheDayByCashier(ctx, req.GetEmployeeId(), day)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &checkpb.GetChecksResponse{Checks: toProtoChecks(checks)}, nil
}

func (h *GRPCHandler) GetChecksOfThePeriodByCashier(ctx context.Context, req *checkpb.GetChecksOfThePeriodByCashierRequest) (*checkpb.GetChecksResponse, error) {
	from, to, err := parseProtoPeriod(req.GetFrom(), req.GetTo())
	if err != nil {
		return nil, err
	}

	checks, err := h.service.GetAllOfThePeriodByCashier(ctx, req.GetEmployeeId(), from, to)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &checkpb.GetChecksResponse{Checks: toProtoChecks(checks)}, nil
}

func (h *GRPCHandler) GetFullCheckData(ctx context.Context, req *checkpb.GetFullCheckDataRequest) (*checkpb.GetFullCheckDataResponse, error) {
	items, err := h.service.GetFullDataByNumber(ctx, req.GetNumber())
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &checkpb.GetFullCheckDataResponse{Items: toProtoFullCheckItems(items)}, nil
}

func (h *GRPCHandler) GetCheckDetailsOfThePeriodByCashier(ctx context.Context, req *checkpb.GetCheckDetailsOfThePeriodByCashierRequest) (*checkpb.GetCheckDetailsResponse, error) {
	from, to, err := parseProtoPeriod(req.GetFrom(), req.GetTo())
	if err != nil {
		return nil, err
	}

	details, err := h.service.GetDetailsOfThePeriodByCashier(ctx, req.GetEmployeeId(), from, to)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &checkpb.GetCheckDetailsResponse{Details: toProtoDetails(details)}, nil
}

func (h *GRPCHandler) GetCheckDetailsOfThePeriod(ctx context.Context, req *checkpb.GetCheckDetailsOfThePeriodRequest) (*checkpb.GetCheckDetailsResponse, error) {
	from, to, err := parseProtoPeriod(req.GetFrom(), req.GetTo())
	if err != nil {
		return nil, err
	}

	details, err := h.service.GetDetailsOfThePeriod(ctx, from, to)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &checkpb.GetCheckDetailsResponse{Details: toProtoDetails(details)}, nil
}

func (h *GRPCHandler) GetSumOfAllChecksOfThePeriodByCashier(ctx context.Context, req *checkpb.GetSumOfAllChecksOfThePeriodByCashierRequest) (*checkpb.CheckTotalSumResponse, error) {
	from, to, err := parseProtoPeriod(req.GetFrom(), req.GetTo())
	if err != nil {
		return nil, err
	}

	total, err := h.service.GetSumOfThePeriodByCashier(ctx, req.GetEmployeeId(), from, to)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &checkpb.CheckTotalSumResponse{TotalSum: total}, nil
}

func (h *GRPCHandler) GetSumOfAllChecksOfThePeriod(ctx context.Context, req *checkpb.GetSumOfAllChecksOfThePeriodRequest) (*checkpb.CheckTotalSumResponse, error) {
	from, to, err := parseProtoPeriod(req.GetFrom(), req.GetTo())
	if err != nil {
		return nil, err
	}

	total, err := h.service.GetSumOfThePeriod(ctx, from, to)
	if err != nil {
		return nil, common.ToStatusError(err)
	}

	return &checkpb.CheckTotalSumResponse{TotalSum: total}, nil
}

func parseProtoPeriod(fromValue, toValue string) (from time.Time, to time.Time, err error) {
	from, err = common.ParseProtoTime(fromValue)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	to, err = common.ParseProtoTime(toValue)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return from, to, nil
}

func toProtoChecks(checks []Check) []*checkpb.Check {
	result := make([]*checkpb.Check, 0, len(checks))
	for _, check := range checks {
		check := check
		result = append(result, toProtoCheck(&check))
	}

	return result
}

func toProtoCheck(check *Check) *checkpb.Check {
	if check == nil {
		return nil
	}

	return &checkpb.Check{
		Number:     check.Number,
		EmployeeId: check.EmployeeID,
		CardNumber: check.CardNumber,
		PrintDate:  common.FormatProtoTime(check.PrintDate),
		SumTotal:   check.SumTotal,
		Vat:        check.VAT,
	}
}

func toProtoFullCheckItems(items []FullCheckData) []*checkpb.FullCheckItem {
	result := make([]*checkpb.FullCheckItem, 0, len(items))
	for _, item := range items {
		result = append(result, &checkpb.FullCheckItem{
			CheckNumber:        item.CheckNumber,
			PrintDate:          common.FormatProtoTime(item.PrintDate),
			SumTotal:           item.SumTotal,
			Vat:                item.VAT,
			Upc:                item.UPC,
			Quantity:           int32(item.Quantity),
			SellingPrice:       item.SellingPrice,
			ProductName:        item.ProductName,
			EmployeeSurname:    item.EmployeeSurname,
			EmployeeName:       item.EmployeeName,
			EmployeePatronymic: item.EmployeePatronymic,
		})
	}

	return result
}

func toProtoDetails(details []Detail) []*checkpb.CheckDetail {
	result := make([]*checkpb.CheckDetail, 0, len(details))
	for _, detail := range details {
		result = append(result, &checkpb.CheckDetail{
			CheckNumber:  detail.CheckNumber,
			PrintDate:    common.FormatProtoTime(detail.PrintDate),
			SumTotal:     detail.SumTotal,
			ProductName:  detail.ProductName,
			Upc:          detail.UPC,
			Quantity:     int32(detail.Quantity),
			SellingPrice: detail.SellingPrice,
		})
	}

	return result
}
