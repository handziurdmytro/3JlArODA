package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/handziurdmytro/3JlArODA/auth-service/internal/business"
	"github.com/handziurdmytro/3JlArODA/auth-service/internal/config"
	"github.com/handziurdmytro/3JlArODA/auth-service/internal/crypto"
	"github.com/handziurdmytro/3JlArODA/auth-service/internal/gateway"
	"github.com/handziurdmytro/3JlArODA/auth-service/internal/middleware"
	"github.com/handziurdmytro/3JlArODA/auth-service/internal/repository"
	"github.com/handziurdmytro/3JlArODA/auth-service/pb"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Load()

	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		slog.Error("database is unreachable", slog.String("error", err.Error()))
		os.Exit(1)
	}
	slog.Info("connected to database")

	cryptoClient := crypto.NewClient(cfg.CryptoServiceAddr)
	defer cryptoClient.Close()

	businessClient := business.NewClient(cfg.BusinessServiceAddr)
	defer businessClient.Close()

	repo := repository.New(db)

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", cfg.Port))
	if err != nil {
		slog.Error("failed to listen", slog.String("port", cfg.Port), slog.String("error", err.Error()))
		os.Exit(1)
	}

	err = repo.SeedDefaultAdmin(context.Background(), cfg.DefaultAdminUser, cfg.DefaultAdminPass, cryptoClient.HashPassword)
	if err != nil {
		slog.Error("failed to seed default admin", slog.String("error", err.Error()))
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.LoggingInterceptor),
	)

	gatewayServer := gateway.NewServer(repo, cryptoClient, businessClient)
	pb.RegisterAuthServiceServer(grpcServer, gatewayServer)

	slog.Info("auth-service started", slog.String("port", cfg.Port))
	if err := grpcServer.Serve(listener); err != nil {
		slog.Error("failed to serve", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
