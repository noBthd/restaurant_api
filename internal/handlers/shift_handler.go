package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/models"
	"github.com/noBthd/restaurant_api.git/internal/services"
)

func CreateShiftHandler(c *gin.Context) {
	var shift models.Shift
	if err := c.ShouldBindJSON(&shift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
			"details": err.Error(),
		})
		return
	}

	if err := services.CreateShift(&shift); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create shift",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Shift created successfully", "shift_id": shift.ID})
}