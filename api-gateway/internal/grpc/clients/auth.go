package clients

import (
	"context"
	"fmt"

	"github.com/handziurdmytro/3JlArODA/api-gateway/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Auth struct {
	grpcClient pb.AuthServiceClient
	conn       *grpc.ClientConn
}

func NewAuthClient(address string) (*Auth, error) {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create a grpc client with %s: %w", address, err)
	}

	return &Auth{
		grpcClient: pb.NewAuthServiceClient(conn),
		conn:       conn,
	}, nil
}

func (cl *Auth) Register(ctx context.Context, username, password string) (*pb.RegisterResponse, error) {
	return cl.grpcClient.Register(ctx, &pb.RegisterRequest{
		Username: username,
		Password: password,
	})
}

func (cl *Auth) Login(ctx context.Context, username, password string) (*pb.LoginResponse, error) {
	return cl.grpcClient.Login(ctx, &pb.LoginRequest{
		Username: username,
		Password: password,
	})
}

func (cl *Auth) Validate(ctx context.Context, token string) (*pb.ValidateResponse, error) {
	return cl.grpcClient.Validate(ctx, &pb.ValidateRequest{
		Token: token,
	})
}

func (cl *Auth) Close() error {
	return cl.conn.Close()
}
