package router

import (
	"github.com/Michael-Seth/taskeeper/internal/interfaces/http/handlers"
	usecase "github.com/Michael-Seth/taskeeper/internal/usecases"
	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.RouterGroup, taskUseCase usecase.TaskUseCase) {
	// Initialize the TaskHandler with the TaskUseCase dependency
	taskHandler := handlers.NewTaskHandler(taskUseCase)

	// Group routes for tasks under "/tasks"
	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.POST("/", taskHandler.CreateTask) // Route to create a task
		taskRoutes.GET("/", taskHandler.GetAllTasks) // Route to fetch all tasks
		// Add more task-related routes here
	}
}
