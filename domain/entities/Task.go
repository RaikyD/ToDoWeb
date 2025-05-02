package entities

import (
	"github.com/google/uuid"
	"time"
)

// TaskPriorityValue represents the priority level of a Task.
type TaskPriorityValue int

const (
	TaskPriorityLow TaskPriorityValue = iota
	TaskPriorityMedium
	TaskPriorityHigh
)

// Task represents a user's goal or todo item with metadata and state.
type Task struct {
	ID          uuid.UUID         // unique task identifier
	UserID      uuid.UUID         // owner user ID
	Name        string            // title of the task
	Description *string           // optional detailed description
	Deadline    *time.Time        // optional due date
	CreatedAt   time.Time         // timestamp when task was created
	Priority    TaskPriorityValue // priority level
	Done        bool              // completion state
}

// NewTask creates a new Task with validation of name and deadline.
// Returns ErrEmptyTaskName if name is empty, ErrDeadlineInPast if deadline is in the past.
func NewTask(name string, userID uuid.UUID, description *string, deadline *time.Time, priority TaskPriorityValue) (*Task, error) {
	if name == "" {
		return nil, ErrEmptyTaskName
	}
	now := time.Now()
	if deadline != nil && deadline.Before(now) {
		return nil, ErrDeadlineInPast
	}
	return &Task{
		ID:          uuid.New(),
		UserID:      userID,
		Name:        name,
		Description: description,
		Deadline:    deadline,
		CreatedAt:   now,
		Priority:    priority,
		Done:        false,
	}, nil
}

// MarkDone marks the task as completed. Returns ErrAlreadyDone if already done.
func (t *Task) MarkDone() error {
	if t.Done {
		return ErrAlreadyDone
	}
	t.Done = true
	return nil
}

// UpdateDeadline changes the task's deadline if provided and valid.
// Passing nil clears the deadline. Returns ErrDeadlineInPast for invalid dates.
func (t *Task) UpdateDeadline(deadline *time.Time) error {
	if deadline == nil {
		t.Deadline = nil
		return nil
	}
	if deadline.Before(time.Now()) {
		return ErrDeadlineInPast
	}
	t.Deadline = deadline
	return nil
}

// UpdateDescription sets or clears the task's description.
func (t *Task) UpdateDescription(desc *string) {
	t.Description = desc
}

// ChangePriority updates the task's priority level.
func (t *Task) ChangePriority(p TaskPriorityValue) {
	t.Priority = p
}
