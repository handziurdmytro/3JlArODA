package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	CryptoServiceAddr   string
	BusinessServiceAddr string
	DatabaseURL         string
	DefaultAdminUser    string
	DefaultAdminPass    string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env file not found, reading from environment")
	}

	return &Config{
		Port:                getOrDefault("PORT", "3131"),
		CryptoServiceAddr:   getOrDefault("CRYPTO_SERVICE_ADDR", "crypto-service:3030"),
		BusinessServiceAddr: getOrDefault("BUSINESS_SERVICE_ADDR", "business-service:2433"),
		DatabaseURL:         getRequired("DATABASE_URL"),
		DefaultAdminUser:    getOrDefault("DEFAULT_ADMIN_USER", "zlahoda@ukma.edu.ua"),
		DefaultAdminPass:    getOrDefault("DEFAULT_ADMIN_PASS", "secret secret"),
	}
}

func getOrDefault(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	slog.Warn("environment variable not set, using default",
		slog.String("key", key),
		slog.String("default", defaultVal),
	)
	return defaultVal
}

func getRequired(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	slog.Error("required config is not set", slog.String("key", key))
	os.Exit(1)
	return ""
}
