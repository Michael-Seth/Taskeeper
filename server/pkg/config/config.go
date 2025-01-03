package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App App `env-required:"true"`
	PG  PG  `env-required:"true"`
}

// App - configuration related to the app itself.
type App struct {
	Name    string
	Version string
}

// PG - PostgreSQL database configuration.
type PG struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

// LoadEnv loads environment variables from the .env file if present.
func LoadEnv() {
	envFilePath := ".env" // One level up from cmd/ folder to reach server/.env
	if err := godotenv.Load(envFilePath); err != nil {
		log.Println("No .env file found, using system environment variables")
		log.Printf("Error loading .env file from %s: %v\n", envFilePath, err)
	}
}

// NewConfig initializes the config struct, reading from environment variables.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	// Load environment variables
	LoadEnv()

	// Populate the struct from environment variables
	cfg.App.Name = os.Getenv("APP_NAME")
	cfg.App.Version = os.Getenv("APP_VERSION")
	cfg.PG.Host = os.Getenv("DB_HOST")
	cfg.PG.User = os.Getenv("DB_USER")
	cfg.PG.Password = os.Getenv("DB_PASSWORD")
	cfg.PG.Database = os.Getenv("DB_NAME")
	cfg.PG.Port = os.Getenv("DB_PORT")

	// Check for missing variables and log them
	missingVars := []string{}
	if cfg.App.Name == "" {
		missingVars = append(missingVars, "APP_NAME")
	}
	if cfg.App.Version == "" {
		missingVars = append(missingVars, "APP_VERSION")
	}
	if cfg.PG.Host == "" {
		missingVars = append(missingVars, "DB_HOST")
	}
	if cfg.PG.User == "" {
		missingVars = append(missingVars, "DB_USER")
	}
	if cfg.PG.Password == "" {
		missingVars = append(missingVars, "DB_PASSWORD")
	}
	if cfg.PG.Database == "" {
		missingVars = append(missingVars, "DB_NAME")
	}
	if cfg.PG.Port == "" {
		missingVars = append(missingVars, "DB_PORT")
	}

	// If any variables are missing, return an error
	if len(missingVars) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	return cfg, nil
}
