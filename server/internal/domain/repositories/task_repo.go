package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Michael-Seth/taskeeper/internal/domain/entities"
)

// TaskRepository defines the interface for task-related database operations.
type TaskRepository interface {
	CreateTask(ctx context.Context, task *entities.Task) error
	GetAllTasks(ctx context.Context) ([]entities.Task, error)
}

// taskRepository is the concrete implementation of TaskRepository.
type taskRepository struct {
	db *sql.DB
}

// NewTaskRepository creates a new instance of taskRepository.
func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

// CreateTask inserts a new task into the database.
func (r *taskRepository) CreateTask(ctx context.Context, task *entities.Task) error {
	query := `INSERT INTO tasks (title, description, due_date, status) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, task.Title, task.Description, task.DueDate, task.Status).Scan(&task.ID)
	if err != nil {
		return errors.New("failed to create task: " + err.Error())
	}
	return nil
}

// GetAllTasks retrieves all tasks from the database.
func (r *taskRepository) GetAllTasks(ctx context.Context) ([]entities.Task, error) {
	query := `SELECT id, title, description, due_date, status FROM tasks`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.New("failed to fetch tasks: " + err.Error())
	}
	defer rows.Close()

	var tasks []entities.Task
	for rows.Next() {
		var task entities.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status); err != nil {
			return nil, errors.New("failed to scan task: " + err.Error())
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New("error iterating over tasks: " + err.Error())
	}

	return tasks, nil
}
