package config

import (
	"log"
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
		log.Println("[WARN] .env file not found, reading from environment")
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

	log.Printf("[WARN] %s not set, using default: %s\n", key, defaultVal)
	return defaultVal
}

func getRequired(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	log.Fatalf("[FATAL] required config %s is not set", key)
	return ""
}

func getDuration(key, defaultVal string) time.Duration {
	val := getOrDefault(key, defaultVal)
	duration, err := time.ParseDuration(val)
	if err != nil {
		log.Fatalf("[FATAL] %s must be a valid duration (e.g. 24h, 30m): %v", key, err)
	}
	return duration
}
