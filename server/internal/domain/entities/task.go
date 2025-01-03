package entities

import "time"

// Task represents a task in the system
// @Description A task with a title, description, completion status, and due date
// @Accept json
// @Produce json
type Task struct {
	ID          int       `db:"id" json:"id"`                   // @Description The unique identifier of the task.
	Title       string    `db:"title" json:"title"`             // @Description The title of the task.
	Description string    `db:"description" json:"description"` // @Description A brief description of the task.
	Completed   bool      `db:"completed" json:"completed"`     // @Description Status if the task is completed.
	DueDate     time.Time `db:"due_date" json:"due_date"`       // @Description The due date of the task.
	Status      string    `db:"status" json:"status"`           // @Description The status of the task.
	UserID      int       `db:"user_id" json:"user_id"`         // @Description The ID of the user associated with the task.
}
