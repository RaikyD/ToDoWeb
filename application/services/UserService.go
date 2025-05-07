package services

import (
	"ToDoWeb/application/dto"
	"ToDoWeb/domain/Interfaces"
	"ToDoWeb/domain/entities"
	"context"
	"github.com/google/uuid"
)

// UserService orchestrates user-related use cases.
type UserService struct {
	userRepo Interfaces.IUserRepository
}

// NewUserService creates a new UserService with the provided repository.
func NewUserService(repo Interfaces.IUserRepository) *UserService {
	return &UserService{userRepo: repo}
}

// CreateUser registers a new user and returns its data.
func (s *UserService) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	user, err := entities.NewUser(req.Name, req.Password)
	if err != nil {
		return nil, err
	}
	user, err = s.userRepo.CreateUser(ctx, user.Name, user.Password)
	if err != nil {
		return nil, err
	}
	return dto.NewUserResponse(user), nil
}

// GetUser retrieves a user by ID.
func (s *UserService) GetUser(ctx context.Context, userID uuid.UUID) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dto.NewUserResponse(user), nil
}

// UpdateUser modifies user fields and returns the updated data.
func (s *UserService) UpdateUser(ctx context.Context, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetUser(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Password != nil {
		user.Password = *req.Password
	}
	if req.Blocked != nil {
		if *req.Blocked {
			if err := user.Block(); err != nil {
				return nil, err
			}
		} else {
			if err := user.Unblock(); err != nil {
				return nil, err
			}
		}
	}
	if err := s.userRepo.Update(ctx, user.ID, *user); err != nil {
		return nil, err
	}
	return dto.NewUserResponse(user), nil
}

// DeleteUser removes a user if business rules allow.
func (s *UserService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	user, err := s.userRepo.GetUser(ctx, userID)
	if err != nil {
		return err
	}
	if err := user.CanDelete(); err != nil {
		return err
	}
	return s.userRepo.DeleteUser(ctx, userID)
}
