package database

import (
	"log"

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

func GetAllContests(db *gorm.DB) *gorm.DB {
	return db.Preload(clause.Associations)
}

func GetContestById(db *gorm.DB, id string) *gorm.DB {
	return db.Preload(clause.Associations).Where("contests.id = ?", id)
}

func GetRankByContestID(db *gorm.DB, id string) *gorm.DB {
	ranks := new([]model.Rank)
	// TODO: add query statement for selecting all ranks of a contest, and order result by rank
	if err := db.Table("Submissions").Where("contest_id = ?", id).Find(ranks).Error; err != nil {
		log.Println("[gorm] error while fetching contest rank")
		return nil
	}
	return db
}

func GetRankByProblemID(db *gorm.DB, id string) *gorm.DB {
	// TODO: add query for fetching rank table of a problem
	return db
}

func GetSubmissionsByUserID(db *gorm.DB, id string) *gorm.DB {
	// TODO: query a user's all submissions
	return db
}

// TODO: query a user's history submissions of a problem
func GetSubmissionsOfUserInOneProblem(c *gin.Context) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
