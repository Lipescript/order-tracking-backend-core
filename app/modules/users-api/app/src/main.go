package main

import (
	"github.com/gin-gonic/gin"

	"users-api/app/src/config"
	"users-api/app/src/constants"
	"users-api/app/src/controller"
	"users-api/app/src/database"
)

func main() {

	// Initialize repository, service, and controller
	userRepository := database.CreateNewMongoDbRepo(config.NewMongoConfig())
	userController := controller.NewUserController(*userRepository)

	initServerAndRoutes(userController)

}

func initServerAndRoutes(uc *controller.UserController) {
	router := gin.Default()

	api := router.Group(constants.BaseAPI)
	{
		api.POST("/users", uc.CreateUserDelivery)
		api.GET("/users", uc.GetUserDeliveries)
	}

	router.Run(constants.DefaultLocalPort)
}
