package service

import "users-api/app/src/database"

type UserService struct {
	userRepo database.Repository
}

func (u *UserService) TestRepository() {
	userRepo := database.CreateNewMongoDbRepo()

	userRepo.AddUser()
	userRepo.DisconnectDB()
}

// NewUserService cria uma nova instância de UserService
func NewUserService(userRepo database.Repository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
