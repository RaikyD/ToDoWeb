package services

import (
	"ToDoWeb/application/dto"
	"ToDoWeb/domain/Interfaces"
	"ToDoWeb/domain/entities"
	"context"
	"github.com/google/uuid"
)

// TaskService orchestrates task-related use cases.
type TaskService struct {
	repo Interfaces.ITaskRepository
}

// NewTaskService creates a new TaskService with the provided repository.
func NewTaskService(repo Interfaces.ITaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask adds a new task for the given user.
func (s *TaskService) CreateTask(ctx context.Context, userID uuid.UUID, req *dto.CreateTaskRequest) (*dto.TaskResponse, error) {
	task, err := entities.NewTask(req.Name, userID, req.Description, req.Deadline, entities.TaskPriorityValue(req.Priority))
	if err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, *task); err != nil {
		return nil, err
	}
	return dto.NewTaskResponse(task), nil
}

// GetAllTasks returns all tasks belonging to a user.
func (s *TaskService) GetAllTasks(ctx context.Context, userID uuid.UUID) ([]*dto.TaskResponse, error) {
	tasks, err := s.repo.GetAllByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	var result []*dto.TaskResponse
	for i := range tasks {
		result = append(result, dto.NewTaskResponse(&tasks[i]))
	}
	return result, nil
}

// GetTaskByID retrieves a specific task by its ID and owner.
func (s *TaskService) GetTaskByID(ctx context.Context, userID, taskID uuid.UUID) (*dto.TaskResponse, error) {
	task, err := s.repo.GetById(ctx, userID, taskID)
	if err != nil {
		return nil, err
	}
	return dto.NewTaskResponse(task), nil
}

// UpdateTask applies updates to a task and returns the updated data.
func (s *TaskService) UpdateTask(ctx context.Context, userID, taskID uuid.UUID, req *dto.UpdateTaskRequest) (*dto.TaskResponse, error) {
	task, err := s.repo.GetById(ctx, userID, taskID)
	if err != nil {
		return nil, err
	}
	if req.Name != nil {
		task.Name = *req.Name
	}
	if req.Description != nil {
		task.UpdateDescription(req.Description)
	}
	if req.Deadline != nil {
		if err := task.UpdateDeadline(req.Deadline); err != nil {
			return nil, err
		}
	}
	if req.Priority != nil {
		task.ChangePriority(entities.TaskPriorityValue(*req.Priority))
	}
	if req.Done != nil {
		if *req.Done {
			if err := task.MarkDone(); err != nil {
				return nil, err
			}
		} else {
			task.Done = false
		}
	}
	if err := s.repo.Update(ctx, userID, taskID, *task); err != nil {
		return nil, err
	}
	return dto.NewTaskResponse(task), nil
}

// DeleteTask removes a task if it exists and belongs to the user.
func (s *TaskService) DeleteTask(ctx context.Context, userID, taskID uuid.UUID) error {
	// Ensure task exists
	if _, err := s.repo.GetById(ctx, userID, taskID); err != nil {
		return err
	}
	return s.repo.Delete(ctx, userID, taskID)
}
