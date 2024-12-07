package migrations

import (
	// "fmt"
	"log"
	// "os"
	// "path/filepath"

	"github.com/Michael-Seth/taskeeper/internal/domain/entities"

	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/Michael-Seth/taskeeper/internal/infrastructure/database"
)

func RunMigrations() {
	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Auto-migrate all models
	err = db.AutoMigrate(&entities.Task{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

// func RunMigrations() {
// 	// Get the connection string from environment variables
// 	db, err := database.Connect()
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}
// 	defer db.Close()

// 	// Construct the path to the migrations folder
// 	migrationPath, err := filepath.Abs("server/migrations")
// 	if err != nil {
// 		log.Fatalf("Failed to construct migration path: %v", err)
// 	}

// 	// Initialize the migration
// 	m, err := migrate.New(
// 		"file://"+migrationPath,
// 		os.Getenv("DB_CONN"), // Use your actual connection string environment variable
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to initialize migrations: %v", err)
// 	}

// 	// Apply all up migrations
// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatalf("Failed to apply migrations: %v", err)
// 	}

// 	fmt.Println("Migrations applied successfully!")
// }
