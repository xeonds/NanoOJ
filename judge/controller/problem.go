package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/database/model"
)

func GetProblems(c *gin.Context) {
	var problems []model.Problem
	if err := db.Find(&problems).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch problems"})
		return
	}
	c.JSON(http.StatusOK, problems)
}

func CreateProblem(c *gin.Context) {
	var problem model.Problem
	if err := c.ShouldBindJSON(&problem); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&problem).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create problem"})
		return
	}
	c.JSON(http.StatusOK, problem)
}

func GetProblemByID(c *gin.Context) {
	var problem model.Problem
	if err := db.Where("id = ?", c.Param("id")).First(&problem).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
		return
	}
	c.JSON(http.StatusOK, problem)
}

func UpdateProblem(c *gin.Context) {
	var problem model.Problem
	if err := db.Where("id = ?", c.Param("id")).First(&problem).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
		return
	}
	if err := c.ShouldBindJSON(&problem); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Save(&problem).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update problem"})
		return
	}
	c.JSON(http.StatusOK, problem)
}

func DeleteProblem(c *gin.Context) {
	var problem model.Problem
	if err := db.Where("id = ?", c.Param("id")).First(&problem).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
		return
	}
	if err := db.Delete(&problem).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete problem"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Problem deleted successfully"})
}
