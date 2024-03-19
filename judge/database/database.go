package database

import (
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
