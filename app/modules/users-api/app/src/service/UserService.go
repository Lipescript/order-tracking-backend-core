package service

import "users-api/app/src/database"

type UserService struct {
	userRepo database.Repository
}

// NewUserService cria uma nova instância de UserService
func NewUserService(userRepo database.Repository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
