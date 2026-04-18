package clients

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Business struct {
	conn *grpc.ClientConn

	Product      *ProductClient
	CustomerCard *CustomerCardClient
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
		Product:      NewProductClient(conn),
		CustomerCard: NewCustomerCardClient(conn),
	}, nil
}

func (cl *Business) Close() error {
	return cl.conn.Close()
}
