// users-api/app/src/controller/user_controller.go
package controller

import "users-api/app/src/service"

// UserController define a estrutura do controlador de usuários
type UserController struct {
	userSvc service.UserService
}

// NewUserController cria uma nova instância de UserController
func NewUserController(userSvc service.UserService) *UserController {
	return &UserController{
		userSvc: userSvc,
	}
}
