package main

import (
	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/config"
	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/handlers"
)

func main() {
	// Load configuration
	db.ConnectDB(config.GetConfig())

	// Initialize the Gin router
	router := gin.Default()

	// Define routes
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

	router.GET("/user", handlers.GetUserByEmailHandler)

	router.Run(":8080")	
}