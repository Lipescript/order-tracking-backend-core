package service

import "users-api/app/src/repository"

type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService cria uma nova instância de UserService
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
