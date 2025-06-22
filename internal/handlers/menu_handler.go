package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/services"
)

func GetAllMenuItemsHandler(c *gin.Context) {
	menuItems, err := services.GetAllMenuItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu items"})
		return
	}

	if len(menuItems) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No menu items found"})
		return
	}

	c.JSON(http.StatusOK, menuItems)
}