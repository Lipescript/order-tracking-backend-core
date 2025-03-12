package main

import (
	"fmt"
	"log"
	"net/http"
	"users-api/app/src/config"
	"users-api/app/src/controller"
	"users-api/app/src/repository"
	"users-api/app/src/service"
)

func main() {
	// Initialize dependencies
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(*userService)

	// Initialize the application
	app := config.NewInitialization(userRepo, *userService, *userController)

	// Start the server
	fmt.Println("Server starting on :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
