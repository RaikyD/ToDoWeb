package services

import (
	"ToDoWeb/domain/Interfaces"
	"context"
	"github.com/google/uuid"
)

type TaskService struct {
	repo Interfaces.ITaskRepository
}

func NewTaskService(r Interfaces.ITaskRepository) *TaskService {
	return &TaskService{
		repo: r,
	}
}

func (s *TaskService) Create(ctx context.Context, userId uuid.UUID) error {
	return nil
}
