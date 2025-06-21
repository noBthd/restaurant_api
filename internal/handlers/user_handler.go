package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/models"
	"github.com/noBthd/restaurant_api.git/internal/services"

	"net/http"
)

func GetUserByEmailHandler(c *gin.Context) {
	email := c.Query("email")

	user, err := services.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUserHandler(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
 
	var user = models.User{
		Email: email,
		Password: password,
	}

	err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": user,
	})
}