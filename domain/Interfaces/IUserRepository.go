package Interfaces

import (
	"ToDoWeb/domain/entities"
	"context"
	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, userName, userPassword string) (*entities.User, error)
	DeleteUser(ctx context.Context, userId uuid.UUID) error
	GetUser(ctx context.Context, userId uuid.UUID) (*entities.User, error)
	Update(ctx context.Context, userId uuid.UUID, newData entities.User) error
}
