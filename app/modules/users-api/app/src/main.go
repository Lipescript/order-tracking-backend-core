package main

import (
	"users-api/app/src/config"
	"users-api/app/src/controller"
	"users-api/app/src/repository"
	"users-api/app/src/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize dependencies
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(*userService)

	// Initialize the application
	app := config.NewInitialization(userRepo, *userService, *userController)

	router := gin.Default()

	// Register routes
	api := router.Group("/api")
	{
		api.GET("/users")
		api.POST("/users")
	}

	// Start the server
	router.Run(":8080")
}
