package repositories

import (
	// "context"
	// "database/sql"
	// "errors"
	"gorm.io/gorm"

	"github.com/Michael-Seth/taskeeper/internal/domain/entities"
)

// TaskRepository defines the interface for task-related database operations.
type TaskRepository interface {
	CreateTask(task *entities.Task) error
	GetAllTasks() ([]entities.Task, error)
}

// taskRepository is the concrete implementation of TaskRepository.
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository creates a new instance of taskRepository.
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// CreateTask saves a new task to the database.
func (r *taskRepository) CreateTask(task *entities.Task) error {
	return r.db.Create(task).Error
}

// GetAllTasks retrieves all tasks from the database.
func (r *taskRepository) GetAllTasks() ([]entities.Task, error) {
	var tasks []entities.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}
