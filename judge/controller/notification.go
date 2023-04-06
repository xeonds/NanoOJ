package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/database/model"
)

func GetNotifications(c *gin.Context) {
	var notifications []model.Notification
	database.NanoDB.Find(&notifications)
	c.JSON(http.StatusOK, notifications)
}

func GetNotificationByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	var notification model.Notification
	if err := database.NanoDB.First(&notification, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Notification not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notification})
}

func CreateNotification(c *gin.Context) {
	var input model.Notification
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.CreatedAt = time.Now().UTC()
	if err := database.NanoDB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func UpdateNotification(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}
	var notification model.Notification
	if err := database.NanoDB.First(&notification, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Notification not found"})
		return
	}
	var input model.Notification
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UpdatedAt = time.Now().UTC()
	if err := database.NanoDB.Model(&notification).Updates(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notification})
}

func DeleteNotification(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	var notification model.Notification
	if err := database.NanoDB.First(&notification, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Notification not found"})
		return
	}

	if err := database.NanoDB.Delete(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
