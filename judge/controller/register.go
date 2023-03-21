package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/model"
)

func Register(c *gin.Context) {
	// Parse request body
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	if _, err := database.GetUserByUsername(user.Username); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	// Create user
	if err := database.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})

}
