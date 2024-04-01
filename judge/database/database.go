package database

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"xyz.xeonds/nano-oj/model"
)

func GetSubmissionsByStatus(db *gorm.DB, status model.Status) *gorm.DB {
	return db.Where("status = ?", status)
}

func GetProblems(db *gorm.DB, page, pageSize int) *gorm.DB {
	return db.Offset((page - 1) * pageSize).Limit(pageSize)
}

func GetUserByEmail(db *gorm.DB, email string) *gorm.DB {
	return db.Where("email = ?", email)
}

func GetUserByUsername(db *gorm.DB, username string) *gorm.DB {
	return db.Where("username = ?", username)
}

func CreateSubmission(db *gorm.DB, submission *model.Submission) *gorm.DB {
	submission.Status = model.Pending
	return db.Create(submission)
}

func GetAllContests(db *gorm.DB, c *gin.Context) *gorm.DB {
	return db.Preload(clause.Associations)
}

func GetContestById(db *gorm.DB, c *gin.Context) *gorm.DB {
	id := c.Param("id")
	return db.Preload(clause.Associations).Where("contests.id = ?", id)
}

// params: contest_id
func GetRankByContestID(db *gorm.DB, c *gin.Context) *gorm.DB {
	id := c.Param("contest_id")
	return db.Where("contest_id = ?", id).Order("rank DESC")
}

// params: problem_id
func GetRankByProblemID(db *gorm.DB, c *gin.Context) *gorm.DB {
	id := c.Param("problem_id")
	return db.Where("problem_id = ?", id).Order("rank DESC")
}

func GetRank(db *gorm.DB, c *gin.Context) *gorm.DB {
	return db.Order("rank DESC")
}

func GetSubmissionsByProblemID(db *gorm.DB, c *gin.Context) *gorm.DB {
	id := c.Param("id")
	return db.Where("problem_id = ?", id).Order("created_at DESC")
}

func GetSubmissionsByUserID(db *gorm.DB, c *gin.Context) *gorm.DB {
	id := c.Param("id")
	return db.Where("user_id = ?", id)
}

func GetSubmissionsOfUserInOneProblem(db *gorm.DB, c *gin.Context) *gorm.DB {
	problem_id, user_id := c.Param("problem_id"), c.Param("user_id")
	return db.Where("problem_id = ? AND user_id = ?", problem_id, user_id).Order("created_at DESC")
}

// TODO
func GetProblemByIdWithoutTestCases(db *gorm.DB, c *gin.Context) *gorm.DB {
	return db
}
