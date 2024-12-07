package handlers

import (
	"github.com/Michael-Seth/taskeeper/internal/domain/entities"
	usecase "github.com/Michael-Seth/taskeeper/internal/usecases"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	usecase usecase.TaskUseCase
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task entities.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.CreateTask(&task); err != nil {
		c.JSON(500, gin.H{"error": "Could not create task"})
		return
	}
	c.JSON(201, task)
}
