package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/services"
)

func GetAllBillsHandler(c *gin.Context) {
	bills, err := services.GetAllBills()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve bills",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, bills)
}

func GetBillByReservationIDHandler(c *gin.Context) {
	reservationID, err := strconv.Atoi(c.Param("reservation_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid reservation ID",
			"details": err.Error(),
		})
		return
	}

	bill, err := services.GetBillByReservationID(reservationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve bill",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, bill)
}

func PayBillHandler(c *gin.Context) {
	reservationID, err := strconv.Atoi(c.Param("reservation_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid reservation ID",
			"details": err.Error(),
		})
		return
	}

	err = services.PayBill(reservationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to pay bill",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bill paid successfully"})
}