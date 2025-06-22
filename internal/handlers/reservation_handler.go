package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/models"
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

func CreateReservationHandler(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	err := services.CreateReservation(reservation)
	if err != nil {
		log.Printf("Error creating reservation: %v", err)
		c.JSON(http.StatusInternalServerError, 
			gin.H{
				"error": "Failed to create reservation",
				"details": err.Error(),
			})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reservation created successfully"})
}

func CancelReservationHandler(c *gin.Context) {
	reservationID := c.Param("id")
	if reservationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Reservation ID is required"})
		return
	}

	ID, err := strconv.Atoi(reservationID)
	if err != nil {
		log.Printf("Error converting reservation ID to integer: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID format"})
		return
	}

	err = services.CancelReservation(ID)
	if err != nil {
		log.Printf("Error canceling reservation: %v", err)
		c.JSON(http.StatusInternalServerError, 
			gin.H{
				"error": "Failed to cancel reservation",
				"details": err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation canceled successfully"})
}