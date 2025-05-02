package mocks

import (
	"ToDoWeb/domain/entities"
	"context"
	"errors"
	"github.com/google/uuid"
	"sync"
)

type InMemoryUserRepo struct {
	store []entities.User
	mu    sync.Mutex
}

func (u *InMemoryUserRepo) CreateUser(ctx context.Context, userName, password string) (*entities.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	newUser, err := entities.NewUser(userName, password)
	if err != nil {
		return nil, errors.New("aaaaa")
	}
	u.store = append(u.store, *newUser)
	return newUser, nil
}

func (u *InMemoryUserRepo) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	var index int
	var userToDelete *entities.User
	for i := range u.store {
		if u.store[i].ID == userID {
			userToDelete = &u.store[i]
			index = i
			break
		}
	}

	if userToDelete == nil {
		return errors.New("no user to delete")
	}

	if err := userToDelete.CanDelete(); err != nil {
		return err
	}

	u.store = append(u.store[:index], u.store[index+1:]...)
	return nil
}

func (u *InMemoryUserRepo) GetUser(ctx context.Context, userID uuid.UUID) (*entities.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	for i := range u.store {
		if u.store[i].ID == userID {
			return &u.store[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func (u *InMemoryUserRepo) Update(ctx context.Context, userID uuid.UUID, newData entities.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	for i := range u.store {
		if u.store[i].ID == userID {
			u.store[i] = newData
			return nil
		}
	}
	return errors.New("user to update is not found")
}
