package handlers

import (
	"github.com/Michael-Seth/taskeeper/internal/domain/entities"
	usecase "github.com/Michael-Seth/taskeeper/internal/usecases"
	"github.com/gin-gonic/gin"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	usecase usecase.TaskUseCase
}

// NewTaskHandler creates a new TaskHandler and injects the TaskUseCase
func NewTaskHandler(usecase usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		usecase: usecase,
	}
}

// CreateTask handles the creation of a new task
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task entities.Task

	// Bind JSON input to the task entity
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Use the TaskUseCase to create the task
	if err := h.usecase.CreateTask(&task); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create task: " + err.Error()})
		return
	}

	// Respond with the created task
	c.JSON(201, gin.H{"message": "Task created successfully", "task": task})
}

// GetAllTasks handles fetching all tasks
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	// Use the TaskUseCase to fetch all tasks
	tasks, err := h.usecase.GetAllTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve tasks: " + err.Error()})
		return
	}

	// Respond with the list of tasks
	c.JSON(200, gin.H{"tasks": tasks})
}
