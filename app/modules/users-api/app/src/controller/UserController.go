// users-api/app/src/controller/user_controller.go
package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"users-api/app/src/constants"
	"users-api/app/src/database"
	"users-api/app/src/domain"

)

type UserController struct {
	userRepo database.MongodbRepo
}

func NewUserController(userRepo database.MongodbRepo) *UserController {
	return &UserController{
		userRepo: userRepo,
	}
}

func (uc *UserController) CreateUserDelivery(c *gin.Context) {
	/*
            delivery_status: {
              enum: ['pendente', 'em_transito', 'entregue', 'cancelado']
            },
	*/
	delivery := domain.NewUserDelivery(uuid.New().String(), uuid.New().String(), "entregue", time.Now())

	database.InsertData(uc.userRepo.MongodbClient, constants.DatabaseName, constants.DatabaseTrackingCollection, delivery)

	response := domain.UserCreatedResponse(delivery)
	c.JSON(201, gin.H{"response": response})
}

func (uc *UserController) GetUserDeliveries(c *gin.Context) {
	// Implement POST /users route logic here
	c.JSON(201, gin.H{"message": "GET /users route"})
}
