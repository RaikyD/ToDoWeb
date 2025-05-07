package dto

import (
	"ToDoWeb/domain/entities"
	"github.com/google/uuid"
	"time"
)

// CreateTaskRequest represents the payload to create a new task.
type CreateTaskRequest struct {
	Name        string     `json:"name" binding:"required"`
	Description *string    `json:"description,omitempty"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	Priority    int        `json:"priority" binding:"required"`
}

// UpdateTaskRequest represents the payload to update an existing task.
type UpdateTaskRequest struct {
	ID          uuid.UUID  `json:"id" binding:"required"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	Priority    *int       `json:"priority,omitempty"`
	Done        *bool      `json:"done,omitempty"`
}

// TaskResponse represents the task data returned in API responses.
type TaskResponse struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	Priority    int        `json:"priority"`
	Done        bool       `json:"done"`
}

// NewTaskResponse converts an entities.Task to TaskResponse.
func NewTaskResponse(t *entities.Task) *TaskResponse {
	return &TaskResponse{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Deadline:    t.Deadline,
		CreatedAt:   t.CreatedAt,
		Priority:    int(t.Priority),
		Done:        t.Done,
	}
}
