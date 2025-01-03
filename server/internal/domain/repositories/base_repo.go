package repositories

// BaseRepository aggregates all the individual repositories.
type BaseRepository interface {
	Task() TaskRepository // Add other repositories here as needed
}

// baseRepository is the implementation of BaseRepository.
type baseRepository struct {
	taskRepo TaskRepository
}

// NewBaseRepository creates a new instance of BaseRepository.
func NewBaseRepository(taskRepo TaskRepository) BaseRepository {
	return &baseRepository{taskRepo: taskRepo}
}

// Task returns the TaskRepository.
func (r *baseRepository) Task() TaskRepository {
	return r.taskRepo
}
