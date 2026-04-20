package clients

import (
	"context"
	"time"

	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/models"
	checkpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/check"
	"google.golang.org/grpc"
)

type CheckClient struct {
	client checkpb.CheckServiceClient
}

func NewCheckClient(conn grpc.ClientConnInterface) *CheckClient {
	return &CheckClient{
		client: checkpb.NewCheckServiceClient(conn),
	}
}

func (cl *CheckClient) Create(ctx context.Context, req models.CreateCheckRequest) error {
	_, err := cl.client.CreateCheck(ctx, &checkpb.CreateCheckRequest{
		Number:     req.Number,
		EmployeeId: req.EmployeeID,
		CardNumber: req.CardNumber,
		PrintDate:  req.PrintDate.Format(timeFormat),
		SumTotal:   req.SumTotal,
		Vat:        req.VAT,
	})

	return err
}

func (cl *CheckClient) Delete(ctx context.Context, number string) error {
	_, err := cl.client.DeleteCheck(ctx, &checkpb.DeleteCheckRequest{Number: number})
	return err
}

func (cl *CheckClient) GetAll(ctx context.Context) ([]*checkpb.Check, error) {
	resp, err := cl.client.GetAllChecks(ctx, &checkpb.GetAllChecksRequest{})
	if err != nil {
		return nil, err
	}

	return resp.GetChecks(), nil
}

func (cl *CheckClient) GetOfDayByCashier(ctx context.Context, employeeID string, date time.Time) ([]*checkpb.Check, error) {
	resp, err := cl.client.GetChecksOfTheDayByCashier(ctx, &checkpb.GetChecksOfTheDayByCashierRequest{
		EmployeeId: employeeID,
		Date:       date.Format(timeFormat),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetChecks(), nil
}

func (cl *CheckClient) GetOfPeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) ([]*checkpb.Check, error) {
	resp, err := cl.client.GetChecksOfThePeriodByCashier(ctx, &checkpb.GetChecksOfThePeriodByCashierRequest{
		EmployeeId: employeeID,
		From:       from.Format(timeFormat),
		To:         to.Format(timeFormat),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetChecks(), nil
}

func (cl *CheckClient) GetFullData(ctx context.Context, number string) ([]*checkpb.FullCheckItem, error) {
	resp, err := cl.client.GetFullCheckData(ctx, &checkpb.GetFullCheckDataRequest{Number: number})
	if err != nil {
		return nil, err
	}

	return resp.GetItems(), nil
}

func (cl *CheckClient) GetDetailsOfPeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) ([]*checkpb.CheckDetail, error) {
	resp, err := cl.client.GetCheckDetailsOfThePeriodByCashier(ctx, &checkpb.GetCheckDetailsOfThePeriodByCashierRequest{
		EmployeeId: employeeID,
		From:       from.Format(timeFormat),
		To:         to.Format(timeFormat),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetDetails(), nil
}

func (cl *CheckClient) GetDetailsOfPeriod(ctx context.Context, from, to time.Time) ([]*checkpb.CheckDetail, error) {
	resp, err := cl.client.GetCheckDetailsOfThePeriod(ctx, &checkpb.GetCheckDetailsOfThePeriodRequest{
		From: from.Format(timeFormat),
		To:   to.Format(timeFormat),
	})
	if err != nil {
		return nil, err
	}

	return resp.GetDetails(), nil
}

func (cl *CheckClient) GetSumOfPeriodByCashier(ctx context.Context, employeeID string, from, to time.Time) (float64, error) {
	resp, err := cl.client.GetSumOfAllChecksOfThePeriodByCashier(ctx, &checkpb.GetSumOfAllChecksOfThePeriodByCashierRequest{
		EmployeeId: employeeID,
		From:       from.Format(timeFormat),
		To:         to.Format(timeFormat),
	})
	if err != nil {
		return 0, err
	}

	return resp.GetTotalSum(), nil
}

func (cl *CheckClient) GetSumOfPeriod(ctx context.Context, from, to time.Time) (float64, error) {
	resp, err := cl.client.GetSumOfAllChecksOfThePeriod(ctx, &checkpb.GetSumOfAllChecksOfThePeriodRequest{
		From: from.Format(timeFormat),
		To:   to.Format(timeFormat),
	})
	if err != nil {
		return 0, err
	}

	return resp.GetTotalSum(), nil
}
