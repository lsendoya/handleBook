package main

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}

func validateEnvironments() error {
	if err := loadEnvironmentVariables(); err != nil {
		return err
	}

	requiredEnvVars := []string{"SERVER_PORT", "ALLOWED_ORIGINS", "ALLOWED_METHODS",
		"JWT_SECRET_KEY", "DB_HOST", "DB_NAME", "DB_USER",
		"DB_PASSWORD", "DB_PORT", "DB_SSL_MODE"}

	for _, envVar := range requiredEnvVars {
		if Config(envVar) == "" {
			return errors.New("the " + envVar + " env is mandatory")
		}
	}

	return nil
}

func loadEnvironmentVariables() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}
	return nil
}
