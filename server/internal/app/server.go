package app

import (
	"context"
	"log"
	"os"

	"github.com/Michael-Seth/taskeeper/internal/infrastructure/database"
	"github.com/Michael-Seth/taskeeper/internal/seeds"
	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) error {
	// Connect to the database (without saving the db variable)
	_, err := database.Connect()
	if err != nil {
		return err
	}

	if len(os.Args) > 1 && os.Args[1] == "seed" {
		// Seed tasks
		if err := seeds.SeedTasks(); err != nil {
			log.Fatalf("Error seeding tasks: %v", err)
			return err // Return error here if seeding fails
		}
		log.Println("Seeding completed!")
		return nil // Return nil after successful seeding
	}

	// Initialize the Gin router
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// Channel to listen for server errors
	serverErr := make(chan error, 1)

	// Start the server in a goroutine
	go func() {
		log.Println("Server is running on port 8080")
		serverErr <- router.Run(":8080") // Start the server and send any errors to the channel
	}()

	// Wait for either context cancellation or server error
	select {
	case <-ctx.Done():
		log.Println("Shutting down the application...")
		return ctx.Err()
	case err := <-serverErr:
		log.Printf("Server error: %v", err)
		return err
	}
}
