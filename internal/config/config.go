package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppHost    string
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	JWTSecret  string
	JWTExpired time.Duration
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	var tokenExpired time.Duration
	if os.Getenv("TOKEN_EXPIRED") != "" {
		expiration, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRED"))
		if err != nil {
			return nil
		}
		tokenExpired = expiration
	} else {
		tokenExpired = time.Duration(time.Hour * 1)
	}

	return &Config{
		AppHost:    getEnvOrDefault("APP_HOST", "localhost"),
		AppPort:    getEnvOrDefault("APP_PORT", "localhost"),
		DBHost:     getEnvOrDefault("DB_HOST", "localhost"),
		DBPort:     getEnvOrDefault("DB_PORT", "5432"),
		DBUser:     getEnvOrDefault("DB_USER", "postgres"),
		DBPassword: getEnvOrDefault("DB_PASSWORD", ""),
		DBName:     getEnvOrDefault("DB_NAME", "reseller_management"),
		DBSSLMode:  getEnvOrDefault("DB_SSL_MODE", "disable"),
		JWTSecret:  getEnvOrDefault("TOKEN_SECRET", "your-secret-key"),
		JWTExpired: tokenExpired,
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
