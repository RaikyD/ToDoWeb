package Interfaces

import (
	"ToDoWeb/domain/entities"
	"context"
	"github.com/google/uuid"
)

// ITaskRepository defines CRUD operations for Task entities.
type ITaskRepository interface {
	// Create inserts a new Task into the repository.
	Create(ctx context.Context, t entities.Task) error

	// GetAllByUser returns all tasks belonging to the given userID.
	GetAllByUser(ctx context.Context, userID uuid.UUID) ([]entities.Task, error)

	// GetById retrieves a specific task by its ID and owner userID.
	GetById(ctx context.Context, userID, taskID uuid.UUID) (*entities.Task, error)

	// Update modifies an existing task identified by userID and taskID.
	Update(ctx context.Context, userID, taskID uuid.UUID, updated entities.Task) error

	// Delete removes the task identified by userID and taskID from the repository.
	Delete(ctx context.Context, userID, taskID uuid.UUID) error
}
