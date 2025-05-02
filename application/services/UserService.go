package services

import (
	"ToDoWeb/domain/Interfaces"
)

type UserService struct {
	userRepo Interfaces.IUserRepository
}

func NewUserService(u Interfaces.IUserRepository) *UserService {
	return &UserService{
		userRepo: u,
	}
}
