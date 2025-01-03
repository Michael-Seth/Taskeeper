package router

import (
	"github.com/Michael-Seth/taskeeper/internal/domain/repositories"
	"github.com/Michael-Seth/taskeeper/internal/interfaces/http/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all API routes with the given BaseRepository.
// @Summary Setup API routes for tasks
// @Description Defines routes for creating and fetching tasks
// @Tags tasks
// @Param api body gin.RouterGroup true "API Routes"
// @Router /api/v1/tasks/ [post]
// @Router /api/v1/tasks/ [get]
func SetupRoutes(api *gin.RouterGroup, repo repositories.BaseRepository) {
	// Initialize the TaskHandler with the TaskRepository
	taskHandler := handlers.NewTaskHandler(repo.Task()) // Use TaskRepository from BaseRepository

	// Group routes for tasks under "/tasks"
	taskRoutes := api.Group("/tasks")
	{
		// @Summary Create a task
		// @Description Creates a new task based on the provided data
		// @Tags tasks
		// @Accept json
		// @Produce json
		// @Param task body entities.Task true "Task data"
		// @Success 201 {object} entities.Task
		// @Failure 400 {object} map[string]string
		// @Failure 500 {object} map[string]string
		taskRoutes.POST("/", taskHandler.CreateTask) // Route to create a task

		// @Summary Get all tasks
		// @Description Retrieves all tasks from the database
		// @Tags tasks
		// @Produce json
		// @Success 200 {array} entities.Task
		// @Failure 500 {object} map[string]string
		taskRoutes.GET("/", taskHandler.GetAllTasks) // Route to fetch all tasks
	}
}
