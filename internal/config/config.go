package config

import (
	"errors"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppHost      string
	AppPort      string
	DBHost       string
	DBPort       string
	DBUsername   string
	DBPassword   string
	DBName       string
	TokenSecret  string
	TokenExpired time.Duration
}

// Function load and wrap all environment variable as configuration
func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	cfg.AppHost = os.Getenv("APP_HOST")
	cfg.AppPort = os.Getenv("APP_PORT")
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBPort = os.Getenv("DB_PORT")
	cfg.DBUsername = os.Getenv("DB_USERNAME")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_NAME")

	if os.Getenv("TOKEN_SECRET") != "" {
		cfg.TokenSecret = os.Getenv("TOKEN_SECRET")
	} else {
		return nil, errors.New("token secret can't be empty")
	}

	if os.Getenv("TOKEN_EXPIRED") != "" {
		expiration, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRED"))
		if err != nil {
			return nil, err
		}

		cfg.TokenExpired = expiration

	} else {
		return nil, errors.New("token expired can't be empty")
	}

	return cfg, nil
}
