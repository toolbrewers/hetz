package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	AppName      string `validate:"required"`
	AppVersion   string `validate:"required"`
	Port         string `validate:"required"`
	BaseURL      string `validate:"required,http_url"`
	JWTKey       string `validate:"required,gte=32,lte=64"`
	SecureCookie bool
}

func New() *Config {
	secureCookie, err := strconv.ParseBool(os.Getenv("SECURE_COOKIE"))
	if err != nil {
		panic(fmt.Errorf("failed to parse SECURE_COOKIE to a boolean: %w", err))
	}

	config := &Config{
		AppName:      os.Getenv("APP_NAME"),
		AppVersion:   os.Getenv("APP_VERSION"),
		Port:         os.Getenv("PORT"),
		BaseURL:      os.Getenv("BASE_URL"),
		JWTKey:       os.Getenv("JWT_KEY"),
		SecureCookie: secureCookie,
	}

	if err := validator.New().Struct(config); err != nil {
		panic(fmt.Errorf("failed to validate config struct: %w", err))
	}

	return config
}
