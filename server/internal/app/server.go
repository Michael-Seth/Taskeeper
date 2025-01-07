package app

import (
	"context"
	"log"
	"os"

	// _ "github.com/Michael-Seth/taskeeper/docs"                           // Import Swagger docs package
	"github.com/Michael-Seth/taskeeper/internal/app/router"              // Correct import for router
	"github.com/Michael-Seth/taskeeper/internal/domain/repositories"     // Import repositories
	"github.com/Michael-Seth/taskeeper/internal/infrastructure/database" // Import database connection
	"github.com/Michael-Seth/taskeeper/internal/seeds"                   // Import seeds
	"github.com/gin-gonic/gin"
)

// RunWithRouter starts the application with the provided Gin router.
func RunWithRouter(ctx context.Context, engine *gin.Engine) error {
	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		return err
	}

	// Check if the app was called with the "seed" argument
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		// Seed tasks
		if err := seeds.SeedTasks(); err != nil {
			log.Fatalf("Error seeding tasks: %v", err)
			return err // Return error if seeding fails
		}
		log.Println("Seeding completed!")
		return nil // Return nil after successful seeding
	}

	// Initialize the TaskRepository
	taskRepo := repositories.NewTaskRepository(db)

	// Create the BaseRepository
	baseRepo := repositories.NewBaseRepository(taskRepo)

	// @Summary Test App Health
	// @Description Test health app status
	// @Tags health
	// @Produce json
	// @Success 200 {object} map[string]string
	// @Failure 500 {object} map[string]string
	// @Router /health [get]
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// Set up API routes using BaseRepository
	api := engine.Group("/api/v1")
	router.SetupRoutes(api, baseRepo) // Centralize route setup with BaseRepository

	// Channel to listen for server errors
	serverErr := make(chan error, 1)

	// Start the server in a goroutine
	go func() {
		log.Println("Server is running on port 8080")
		serverErr <- engine.Run(":8080") // Start the server and send any errors to the channel
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
