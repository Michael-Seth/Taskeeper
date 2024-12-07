package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Michael-Seth/taskeeper/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		log.Printf("Received signal: %s. Shutting down gracefully...", sig)
		cancel()
	}()

	if err := app.Run(ctx); err != nil {
		log.Fatalf("Application failed to start: %v", err)
	}

	log.Println("Application exited cleanly.")
}
