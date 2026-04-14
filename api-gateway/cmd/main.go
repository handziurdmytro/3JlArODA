package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/config"
	"github.com/handziurdmytro/3JlArODA/api-gateway/internal/routes"
)

func main() {
	cfg := config.Load()

	gin.SetMode(cfg.GinMode)
	router := gin.Default()

	routes.SetupRoutes(router)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf(
		"[INFO] starting api-gateway on %s (gin_mode=%s, business_service=%s, auth_service=%s)",
		addr,
		cfg.GinMode,
		cfg.BusinessService,
		cfg.AuthService,
	)

	err := router.Run(addr)
	if err != nil {
		log.Fatalf("[FATAL] failed to start api-gateway: %v", err)
	}
}
