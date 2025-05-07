package dto

import (
	"ToDoWeb/domain/entities"
	"github.com/google/uuid"
)

// CreateUserRequest represents the payload to register a new user.
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UpdateUserRequest represents the payload to update an existing user.
type UpdateUserRequest struct {
	ID       uuid.UUID `json:"id" binding:"required"`
	Name     *string   `json:"name,omitempty"`
	Password *string   `json:"password,omitempty"`
	Blocked  *bool     `json:"blocked,omitempty"`
}

// UserResponse represents the user data returned in API responses.
type UserResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Blocked bool      `json:"blocked"`
}

// NewUserResponse converts an entities.User to UserResponse.
func NewUserResponse(u *entities.User) *UserResponse {
	return &UserResponse{
		ID:      u.ID,
		Name:    u.Name,
		Blocked: u.Blocked,
	}
}
