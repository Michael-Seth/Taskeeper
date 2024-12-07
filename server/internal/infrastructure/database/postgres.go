package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes a connection to the PostgreSQL database using environment variables.
func Connect() (*gorm.DB, error) {
	// Using environment variables for the connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),     // DB host (database)
		os.Getenv("DB_USER"),     // DB user (postgres)
		os.Getenv("DB_PASSWORD"), // DB password (postgres)
		os.Getenv("DB_NAME"),     // DB name (taskeeper)
		os.Getenv("DB_PORT"),     // DB port (5432)
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
