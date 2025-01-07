package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Michael-Seth/taskeeper/docs" // Import Swagger docs package
	"github.com/Michael-Seth/taskeeper/internal/app"
)

//	@title			Taskeeper API
//	@version		1.0
//	@description	This is a structured API documentation for Taskeeper.
//	@termsOfService	http://example.com/terms/

//	@contact.name	API Support
//	@contact.url	http://example.com/contact
//	@contact.email	support@example.com

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

// @host		localhost:8080
// @BasePath	/
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		log.Printf("Received signal: %s. Shutting down gracefully...", sig)
		cancel()
	}()

	// Create a Gin router
	r := gin.Default()
	env := os.Getenv("ENV")

	// Set trusted proxies based on environment
	if env == "production" {
		err := r.SetTrustedProxies([]string{"192.168.1.1", "192.168.1.2"}) // Replace with your trusted proxies
		if err != nil {
			log.Fatalf("Error setting trusted proxies: %v", err)
		}
		log.Println("Trusted proxies set for production environment.")
	} else {
		// Disable trusted proxies for non-production environments
		err := r.SetTrustedProxies(nil)
		if err != nil {
			log.Fatalf("Error disabling trusted proxies: %v", err)
		}
		log.Println("Trusted proxies disabled for non-production environment.")
	}

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the application
	if err := app.RunWithRouter(ctx, r); err != nil {
		log.Fatalf("Application failed to start: %v", err)
	}

	log.Println("Application exited cleanly.")
}
