package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/handziurdmytro/3JlArODA/auth-service/internal/config"
	"github.com/handziurdmytro/3JlArODA/auth-service/internal/crypto"
	"github.com/handziurdmytro/3JlArODA/auth-service/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("[FATAL] failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("[FATAL] database is unreachable: %v", err)
	}
	log.Println("[INFO] connected to database")

	cryptoClient := crypto.NewClient(cfg.CryptoServiceAddr)
	defer cryptoClient.Close()

	repo := repository.New(db)

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", cfg.Port))
	if err != nil {
		log.Fatalf("[FATAL] failed to listen on port %s: %v", cfg.Port, err)
	}

	grpcServer := grpc.NewServer()

	_ = repo
	_ = cryptoClient

	log.Printf("[INFO] auth-service listening on port %s", cfg.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("[FATAL] failed to serve: %v", err)
	}
}
