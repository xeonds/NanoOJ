package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/database"
	"xyz.xeonds/nano-oj/database/model"
)

func GetContests(c *gin.Context) {
	var contests []model.Contest
	database.NanoDB.Find(&contests)
	c.JSON(http.StatusOK, gin.H{"data": contests})
}

func GetContestByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contest ID"})
		return
	}

	var contest model.Contest
	if err := database.NanoDB.First(&contest, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contest not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contest})
}

func CreateContest(c *gin.Context) {
	var input model.Contest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.NanoDB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contest"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func UpdateContest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contest ID"})
		return
	}

	var contest model.Contest
	if err := database.NanoDB.First(&contest, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contest not found"})
		return
	}

	var input model.Contest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.NanoDB.Model(&contest).Updates(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update contest"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contest})
}

func DeleteContest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contest ID"})
		return
	}

	var contest model.Contest
	if err := database.NanoDB.First(&contest, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contest not found"})
		return
	}

	if err := database.NanoDB.Delete(&contest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contest"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
