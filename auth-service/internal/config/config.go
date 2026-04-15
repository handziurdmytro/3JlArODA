package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	CryptoServiceAddr string
	DatabaseURL       string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env file not found, reading from environment")
	}

	return &Config{
		Port:              getOrDefault("PORT", "2828"),
		CryptoServiceAddr: getOrDefault("CRYPTO_SERVICE_ADDR", "0.0.0.0:2929"),
		DatabaseURL:       getRequired("DATABASE_URL"),
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

func getDuration(key, defaultVal string) time.Duration {
	val := getOrDefault(key, defaultVal)
	duration, err := time.ParseDuration(val)
	if err != nil {
		slog.Error("invalid duration format (e.g. 24h, 30m)",
			slog.String("key", key),
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}
	return duration
}
