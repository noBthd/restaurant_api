package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/models"
	"github.com/noBthd/restaurant_api.git/internal/services"
)

func CreateWaiterHandler(c *gin.Context) {
	var waiter models.Waiter
	if err := c.ShouldBindJSON(&waiter); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
			"details": err.Error(),
		})
		return
	}

	err := services.CreateWaiter(waiter)
	if err != nil {
		log.Printf("Error creating waiter: %v", err)
		c.JSON(http.StatusInternalServerError, 
			gin.H{
				"error": "Failed to create waiter",
				"details": err.Error(),
			})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Waiter created successfully"})
}