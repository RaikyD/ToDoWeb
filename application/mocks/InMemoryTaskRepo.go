package mocks

import (
	"context"
	"errors"
	"sync"

	"ToDoWeb/domain/entities"
	"github.com/google/uuid"
)

// InMemoryTaskRepo is an in-memory implementation of ITaskRepository for testing.
type InMemoryTaskRepo struct {
	mu    sync.RWMutex
	store []entities.Task
}

// NewInMemoryTaskRepo initializes a new in-memory task repository.
func NewInMemoryTaskRepo() *InMemoryTaskRepo {
	return &InMemoryTaskRepo{
		store: make([]entities.Task, 0),
	}
}

// Create adds a new task to the in-memory store.
func (r *InMemoryTaskRepo) Create(ctx context.Context, t entities.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store = append(r.store, t)
	return nil
}

// GetAllByUser returns all tasks for a given userID.
func (r *InMemoryTaskRepo) GetAllByUser(ctx context.Context, userID uuid.UUID) ([]entities.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var tasks []entities.Task
	for _, task := range r.store {
		if task.UserID == userID {
			tasks = append(tasks, task)
		}
	}
	if len(tasks) == 0 {
		return nil, errors.New("task: no tasks found for user")
	}
	return tasks, nil
}

// GetById retrieves a single task by userID and taskID.
func (r *InMemoryTaskRepo) GetById(ctx context.Context, userID, taskID uuid.UUID) (*entities.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i := range r.store {
		if r.store[i].UserID == userID && r.store[i].ID == taskID {
			return &r.store[i], nil
		}
	}
	return nil, errors.New("task: not found")
}

// Update replaces the task identified by userID and taskID with the provided updated task.
func (r *InMemoryTaskRepo) Update(ctx context.Context, userID, taskID uuid.UUID, updated entities.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.store {
		if r.store[i].UserID == userID && r.store[i].ID == taskID {
			r.store[i] = updated
			return nil
		}
	}
	return errors.New("task: no such task to update")
}

// Delete removes the task identified by userID and taskID from the in-memory store.
func (r *InMemoryTaskRepo) Delete(ctx context.Context, userID, taskID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	index := -1
	for i := range r.store {
		if r.store[i].UserID == userID && r.store[i].ID == taskID {
			index = i
			break
		}
	}
	if index < 0 {
		return errors.New("task: no such task to delete")
	}
	r.store = append(r.store[:index], r.store[index+1:]...)
	return nil
}
