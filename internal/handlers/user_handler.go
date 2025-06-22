package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/models"
	"github.com/noBthd/restaurant_api.git/internal/services"

	"net/http"
	"strconv"
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

func LoginUserHandler(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	var user = models.User{
		Email: email,
		Password: password,
	}

	existingUser, err := services.LoginUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to login user",
			"details": err.Error(),
		})
		return
	}

	if existingUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, existingUser)
}

func RemoveUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
			"details": err.Error(),
		})
		return
	}

	err = services.RemoveUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to remove user",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User removed successfully",
		"user_id": id,
	})
}

func MakeUserAdminHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
			"details": err.Error(),
		})
		return
	}

	err = services.MakeUserAdmin(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to make user admin",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User made admin successfully",
		"user_id": id,
	})
}