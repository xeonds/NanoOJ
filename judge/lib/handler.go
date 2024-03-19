package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/model"
)

func HandleReRunJudge(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		submission := new(model.Submission)
		if err := db.First(submission, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
			return
		}
		submission.Status = model.Pending
		if err := db.Save(submission).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Re-run successfully"})
	}
}
