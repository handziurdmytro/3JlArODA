package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	"github.com/handziurdmytro/3JlArODA/business-service/internal/customercard"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/database/postgres"
	"github.com/handziurdmytro/3JlArODA/business-service/internal/product"
	customercardpb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/customercard"
	productpb "github.com/handziurdmytro/3JlArODA/business-service/pb/business/product"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if err := godotenv.Load(); err != nil {
		logger.Info("env file was not loaded", "error", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		logger.Error("DATABASE_URL is required")
		os.Exit(1)
	}

	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "8082"
	}

	pool, err := postgres.GetNewPool(context.Background(), dsn)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	productRepo := product.NewProductRepository(pool)
	productService := product.NewService(productRepo)
	productHandler := product.NewGRPCHandler(productService)

	customerCardRepo := customercard.NewCustomerCardRepository(pool)
	customerCardService := customercard.NewService(customerCardRepo)
	customerCardHandler := customercard.NewGRPCHandler(customerCardService)

	grpcServer := grpc.NewServer()
	productpb.RegisterProductServiceServer(grpcServer, productHandler)
	customercardpb.RegisterCustomerCardServiceServer(grpcServer, customerCardHandler)

	address := ":" + port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Error("failed to listen", "address", address, "error", err)
		os.Exit(1)
	}

	logger.Info("business-service grpc server started", "address", address)
	if err := grpcServer.Serve(listener); err != nil {
		logger.Error("grpc server stopped with error", "error", err)
		os.Exit(1)
	}
}
