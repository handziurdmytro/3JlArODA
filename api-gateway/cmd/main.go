package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/config"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/grpc/clients"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/handlers"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/routes"
)

func main() {
	cfg := config.Load()

	authClient, err := clients.NewAuthClient(cfg.AuthService)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := authClient.Close(); err != nil {
			log.Printf("[WARN] failed to close auth grpc client: %v", err)
		}
	}()

	businessClient, err := clients.NewBusinessClient(cfg.BusinessService)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := businessClient.Close(); err != nil {
			log.Printf("[WARN] failed to close business grpc client: %v", err)
		}
	}()

	authHandler := handlers.NewAuthHandler(authClient)

	gin.SetMode(cfg.GinMode)
	router := gin.Default()

	routes.SetupRoutes(router, authHandler, authClient)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf(
		"[INFO] starting api-gateway on %s (gin_mode=%s, business_service=%s, auth_service=%s)",
		addr,
		cfg.GinMode,
		cfg.BusinessService,
		cfg.AuthService,
	)

	if err := router.Run(addr); err != nil {
		log.Fatalf("[FATAL] failed to start api-gateway: %v", err)
	}
}
