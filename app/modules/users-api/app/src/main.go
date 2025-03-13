package main

import "github.com/gin-gonic/gin"

func main() {

	// Register routes
	router := gin.Default()
	api := router.Group("/api")
	{
		// User routes{
		api.GET("/users")
		api.POST("/users")
	}

	// Start the server
	router.Run(":8080")
}
