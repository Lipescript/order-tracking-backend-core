// users-api/app/src/controller/user_controller.go
package controller

import (
    "fmt"
    "net/http"

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

api := router.Group("/api")

{
    api.GET("/users", userCtrl.GetUsers)
    api.POST("/users", userCtrl.CreateUser)
}

func (uc *UserController) GetUsers(c *gin.Context) {
    fmt.Println("GET /users route")
    c.JSON(http.StatusOK, gin.H{"message": "GET /users route"})
}

// POST /users
func (uc *UserController) CreateUser(c *gin.Context) {
    fmt.Println("POST /users route")
    c.JSON(http.StatusCreated, gin.H{"message": "POST /users route"})
}