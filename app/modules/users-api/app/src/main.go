package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	//"github.com/mrkouhadi/go-graphql-mongo/internal/db"
	//	"github.com/mrkouhadi/go-graphql-mongo/internal/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"users-api/app/src/constants"
	"users-api/app/src/database"
)

func main() {

	initServerAndRoutes()
	initDatabaseConection()

	Mongodb_Repo := database.CreateNewMongoDbRepo(&app)

	// set up handlers
	repo := database.NewRepo(&app, Mongodb_Repo)
	database.NewHandlers(repo)

	// Start the server
}

func initServerAndRoutes() {

	router := gin.Default()
	api := router.Group(constants.BaseAPI)
	{
		// User routes{
		api.GET("/users")
		api.POST("/users")
	}
	router.Run(constants.DefaultLocalPort)

}

func initDatabaseConection() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(constants.DatabaseURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// CHECK CONNECTION
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" --------------- App has been Connected to MongoDB!")
	return client
}
