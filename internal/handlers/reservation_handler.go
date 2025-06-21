package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/services"
)

func GetAllReservationsHandler(c *gin.Context) {
	reservations, err := services.GetAllReservations()
	if err != nil {
		log.Printf("Error fetching reservations: %v", err)
		c.JSON(http.StatusInternalServerError, 
			gin.H{
				"error": "Failed to fetch reservations",
				"details": err.Error(),
			})
		return
	}

	if len(reservations) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No reservations found"})
		return
	}

	c.JSON(http.StatusOK, reservations)
}
