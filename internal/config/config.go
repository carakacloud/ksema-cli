package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	KsemaServerIP string `validate:"required" env:"KSEMA_HOST"`
	KsemaAPIKey   string `validate:"required" env:"KSEMA_API_KEY"`
	KsemaPIN      string
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	godotenv.Load()

	config := &Config{
		KsemaServerIP: os.Getenv("KSEMA_HOST"),
		KsemaAPIKey:   os.Getenv("KSEMA_API_KEY"),
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return config, nil
}
