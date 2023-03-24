package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/database/model"
)

func GetSubmissions(c *gin.Context) {
	var submissions []model.Submission
	if err := database.NanoDB.Find(&submissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, submissions)
}

func CreateSubmission(c *gin.Context) {
	var submission model.Submission
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.NanoDB.Create(&submission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, submission)
}

func GetSubmissionByID(c *gin.Context) {
	var submission model.Submission
	if err := database.NanoDB.Where("id = ?", c.Param("id")).First(&submission).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, submission)
}

func UpdateSubmission(c *gin.Context) {
	var submission model.Submission
	if err := database.NanoDB.Where("id = ?", c.Param("id")).First(&submission).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.NanoDB.Save(&submission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, submission)
}
