package business

import (
	"context"
	"log/slog"
	"os"
	"time"

	rolepb "github.com/handziurdmytro/3JlArODA/auth-service/pb/business/role"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	grpcClient rolepb.RoleServiceClient
	conn       *grpc.ClientConn
}

func NewClient(address string) *Client {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		slog.Error("failed to connect to business-service",
			slog.String("address", address),
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	return &Client{
		grpcClient: rolepb.NewRoleServiceClient(conn),
		conn:       conn,
	}
}

func (c *Client) GetEmployeeRole(ctx context.Context, employeeID string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	resp, err := c.grpcClient.GetEmployeeRole(ctx, &rolepb.GetEmployeeRoleRequest{
		EmployeeId: employeeID,
	})

	if err != nil {
		return "", err
	}

	return resp.Role, nil
}

func (c *Client) Close() {
	if err := c.conn.Close(); err != nil {
		slog.Warn("error closing business-service connection", slog.String("error", err.Error()))
	}
}
