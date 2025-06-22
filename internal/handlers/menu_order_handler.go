package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/models"
	"github.com/noBthd/restaurant_api.git/internal/services"
)

func GetAllMenuOrdersHandler(c *gin.Context) {
	menuOrders, err := services.GetAllMenuOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch menu orders",
			"details": err.Error(),
		})
		return
	}
	
	if len(menuOrders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No menu orders found"})
		return
	}
	c.JSON(http.StatusOK, menuOrders)
}

func GetMenuOrderByIDHandler(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil || orderID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid order ID",
			"details": err.Error(),
		})
		return
	}

	menuOrder, err := services.GetMenuOrderByID(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch menu order",
			"details": err.Error(),
		})
		return
	}

	if len(menuOrder) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Menu order not found"})
		return
	}
	c.JSON(http.StatusOK, menuOrder)
}

func CreateMenuOrderHandler(c *gin.Context) {
	// Parse the request body to get the MenuOrder details
	var menuOrder models.MenuOrder
	if err := c.ShouldBindJSON(&menuOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Get the reservation ID from the query parameters
	reservationID, err := strconv.Atoi(c.Query("reservation_id"))
	if err != nil || reservationID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid reservation ID",
			"details": err.Error(),
		})
		return
	}

	err = services.CreateMenuOrder(&menuOrder, reservationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create menu order",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Menu order created successfully", "order_id": menuOrder.OrderID})
}