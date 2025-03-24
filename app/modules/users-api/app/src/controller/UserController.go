// users-api/app/src/controller/user_controller.go
package controller

import (
	"github.com/gin-gonic/gin"

	"users-api/app/src/service"
)

type UserController struct {
	userSvc service.UserService
}

func NewUserController(userSvc service.UserService) *UserController {
	return &UserController{
		userSvc: userSvc,
	}
}

func getUsers(us *service.UserService, context *gin.Context) {
	// Implement GET /users route logic here

	us.TestRepository()

	context.JSON(200, gin.H{"message": "GET /users route"})
}

func createUser(context *gin.Context) {
	// Implement POST /users route logic here
	context.JSON(201, gin.H{"message": "POST /users route"})
}
