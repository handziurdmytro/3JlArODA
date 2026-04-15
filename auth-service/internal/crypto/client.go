package crypto

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/handziurdmytro/3JlArODA/auth-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	grpcClient pb.AuthSecretServiceClient
	conn       *grpc.ClientConn
}

func NewClient(address string) *Client {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		slog.Error("failed to connect to crypto-service",
			slog.String("address", address),
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	c := &Client{
		grpcClient: pb.NewAuthSecretServiceClient(conn),
		conn:       conn,
	}

	c.ping()

	return c
}

func (c *Client) ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.grpcClient.HashPassword(ctx, &pb.HashRequest{PlainPassword: "ping"})
	if err != nil {
		slog.Error("crypto-service is unreachable", slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Info("connected to crypto-service successfully")
}

func (c *Client) HashPassword(plainPassword string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.grpcClient.HashPassword(ctx, &pb.HashRequest{PlainPassword: plainPassword})
	if err != nil {
		return "", err
	}

	return resp.HashString, nil
}

func (c *Client) VerifyPassword(plainPassword, hashString string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.grpcClient.VerifyPassword(ctx, &pb.VerifyRequest{
		PlainPassword: plainPassword,
		HashString:    hashString,
	})

	if err != nil {
		return false, err
	}

	return resp.IsValid, nil
}

func (c *Client) SignJWT(userID, username string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.grpcClient.SignJWT(ctx, &pb.SignJWTRequest{
		UserId:   userID,
		Username: username,
	})

	if err != nil {
		return "", err
	}

	return resp.Token, nil
}

func (c *Client) ValidateJWT(token string) (*pb.ValidateJWTResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.grpcClient.ValidateJWT(ctx, &pb.ValidateJWTRequest{
		Token: token,
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Close() {
	if err := c.conn.Close(); err != nil {
		slog.Warn("error closing crypto-service connection", slog.String("error", err.Error()))
	}
}
