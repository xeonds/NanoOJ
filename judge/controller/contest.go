package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/database/model"
)

func GetContests(c *gin.Context) {
	var contests []model.Contest
	db.Find(&contests)
	c.JSON(http.StatusOK, contests)
}

func GetContestByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contest ID"})
		return
	}

	var contest model.Contest
	if err := db.First(&contest, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contest not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contest})
}

func CreateContest(c *gin.Context) {
	var input struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		StartTime   time.Time `json:"start_time"`
		EndTime     time.Time `json:"end_time"`
		Problems    []uint32  `json:"problems"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var problems []model.Problem
	for _, id := range input.Problems {
		problem, err := repo.GetProblemByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		problems = append(problems, *problem)
	}

	contest := model.Contest{
		Title:       input.Title,
		Description: input.Description,
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
		Problems:    problems,
	}

	if err := db.Create(&contest).Error; err != nil {
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
	if err := db.First(&contest, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contest not found"})
		return
	}

	var input model.Contest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&contest).Updates(&input).Error; err != nil {
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
	if err := db.First(&contest, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contest not found"})
		return
	}

	if err := db.Delete(&contest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contest"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
