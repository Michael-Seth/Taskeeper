package database

import (
	"fmt"

	"github.com/Michael-Seth/taskeeper/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes a connection to the PostgreSQL database using environment variables.
func Connect() (*gorm.DB, error) {
	// Load configuration
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	// Using the loaded config to build the connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PG.Host,     // DB host
		cfg.PG.User,     // DB user
		cfg.PG.Password, // DB password
		cfg.PG.Database, // DB name
		cfg.PG.Port,     // DB port
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
