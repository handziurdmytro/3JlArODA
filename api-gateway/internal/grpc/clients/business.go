package clients

import (
	"fmt"

	customercardpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/customercard"
	productpb "github.com/handziurdmytro/3JlArODA/api-gateway/pb/business/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Business struct {
	conn *grpc.ClientConn

	Product      productpb.ProductServiceClient
	CustomerCard customercardpb.CustomerCardServiceClient
}

func NewBusinessClient(address string) (*Business, error) {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create a business grpc client with %s: %w", address, err)
	}

	return &Business{
		conn:         conn,
		Product:      productpb.NewProductServiceClient(conn),
		CustomerCard: customercardpb.NewCustomerCardServiceClient(conn),
	}, nil
}

func (cl *Business) Close() error {
	return cl.conn.Close()
}
