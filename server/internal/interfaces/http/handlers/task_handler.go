package handlers

import (
	"net/http"

	"github.com/Michael-Seth/taskeeper/internal/domain/entities"

	"github.com/Michael-Seth/taskeeper/internal/domain/repositories"
	"github.com/gin-gonic/gin"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	taskRepo repositories.TaskRepository
}

// NewTaskHandler creates a new TaskHandler and injects the TaskUseCase
func NewTaskHandler(taskRepo repositories.TaskRepository) *TaskHandler {
	return &TaskHandler{taskRepo: taskRepo}
}

// CreateTask handles the creation of a new task.
// @Summary Create a new task
// @Description Create a new task with the provided details
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body entities.Task true "Task details"
// @Success 201 {object} entities.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/tasks/ [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task entities.Task

	// Bind JSON input to the task entity.
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Use the TaskRepository to create the task.
	if err := h.taskRepo.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task: " + err.Error()})
		return
	}

	// Respond with the created task.
	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    task,
	})
}

// GetAllTasks handles fetching all tasks.
// @Summary Retrieve all tasks
// @Description Fetch all tasks from the database
// @Tags tasks
// @Produce json
// @Success 200 {array} entities.Task
// @Failure 500 {object} map[string]string
// @Router /api/v1/tasks/ [get]
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	// Use the TaskRepository to fetch all tasks.
	tasks, err := h.taskRepo.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks: " + err.Error()})
		return
	}

	// Respond with the list of tasks.
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
