package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	GinMode         string
	BusinessService string
	AuthService     string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARN] .env file not found, reading from environment")
	}

	return &Config{
		Port:            getOrDefault("PORT", "8080"),
		GinMode:         getOrDefault("GIN_MODE", "debug"),
		BusinessService: getOrDefault("BUSINESS_SERVICE_ADDR", "localhost:8082"),
		AuthService:     getOrDefault("AUTH_SERVICE_ADDR", "localhost:2828"),
	}
}

func getOrDefault(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	log.Printf("[WARN] %s not set, using default: %s\n", key, defaultVal)
	return defaultVal
}
