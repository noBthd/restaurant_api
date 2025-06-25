package handlers

import (
	"net/http"
	"strconv"

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

	tableID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create shift",
			"details": err.Error(),
		})
		return
	}

	if err := services.CreateShift(&shift, tableID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create shift",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Shift created successfully", "shift_id": shift.ID})
}

func AddServedTableHandler(c *gin.Context) {
	ShiftID, err := strconv.Atoi(c.Param("shift_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid shift ID",
			"details": err.Error(),
		})
		return
	}

	if err := services.AddServedTable(ShiftID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add served table",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Served table added successfully"})
}

func AddServedTableByReservationHandler(c *gin.Context) {
	reservationID, err := strconv.Atoi(c.Param("reservation_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid reservation ID",
			"details": err.Error(),
		})
		return
	}

	if err := services.AddServedTableByReservation(reservationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add served table by reservation",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Served table added successfully by reservation"})
}