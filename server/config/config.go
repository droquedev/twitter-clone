package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	DB_URI     string `validate:"required"`
	JWT_SECRET string `validate:"required"`
}

var config *Config

func (c *Config) validate() error {
	validator := validator.New()
	return validator.Struct(c)
}

func NewConfig() (*Config, error) {

	config = &Config{
		DB_URI:     os.Getenv("DB_URI"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}

	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("Missing env vars\n%w", err)
	}

	return config, nil
}
