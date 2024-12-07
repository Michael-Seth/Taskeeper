package seeds

import (
	"log"

	"github.com/Michael-Seth/taskeeper/internal/domain/entities"
	"github.com/Michael-Seth/taskeeper/internal/infrastructure/database"
)

func SeedTasks() error {
	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		return err
	}
	// defer db.Close()

	// Seed data
	tasks := []entities.Task{
		{Title: "Task 1", Description: "This is task 1", Completed: false},
		{Title: "Task 2", Description: "This is task 2", Completed: true},
		{Title: "Task 3", Description: "This is task 3", Completed: false},
	}

	// Insert seed data into the tasks table
	if err := db.Create(&tasks).Error; err != nil {
		return err
	}

	log.Println("Seeded tasks successfully!")
	return nil
}
