package handlers

import (
	"net/http"

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