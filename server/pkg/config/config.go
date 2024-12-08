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
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
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

	// Ensure that all required fields are populated
	if cfg.App.Name == "" || cfg.App.Version == "" || cfg.PG.Host == "" || cfg.PG.User == "" || cfg.PG.Password == "" || cfg.PG.Database == "" || cfg.PG.Port == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return cfg, nil
}
