package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/database/model"
)

func GetSubmissions(c *gin.Context) {
	var submissions []model.Submission
	if err := db.Find(&submissions).Error; err != nil {
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
	problem, err := repo.GetProblemByID(submission.ProblemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	submission.Problem = *problem
	if err := db.Create(&submission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, submission)
}

func GetSubmissionByID(c *gin.Context) {
	var submission model.Submission
	if err := db.Where("id = ?", c.Param("id")).First(&submission).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, submission)
}

func UpdateSubmission(c *gin.Context) {
	var submission model.Submission
	if err := db.Where("id = ?", c.Param("id")).First(&submission).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Save(&submission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, submission)
}
