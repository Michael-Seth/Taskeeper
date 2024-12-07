package usecase

import "github.com/Michael-Seth/taskeeper/internal/domain/entities"

type TaskUseCase interface {
	CreateTask(task *entities.Task) error
	GetAllTasks() ([]entities.Task, error)
}
