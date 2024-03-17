package database

import (
	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/lib"
	"xyz.xeonds/nano-oj/model"
)

// Submissions
func CreateSubmission(db *gorm.DB, submission *model.Submission) *gorm.DB {
	return db.Create(&submission)
}

func GetSubmissionByID(db *gorm.DB, id uint32) *gorm.DB {
	return db.Where("id = ?", id)
}

func GetSubmissionsByProblemID(db *gorm.DB, problemID uint32) *gorm.DB {
	return db.Where("id = ?", problemID)
}

func GetSubmissionsByUserID(db *gorm.DB, userID uint32) *gorm.DB {
	return db.Where("id = ?", userID)
}

func GetSubmissionsByStatus(db *gorm.DB, status model.Status) *gorm.DB {
	return db.Where("status = ?", status)
}

func UpdateSubmission(db *gorm.DB, submission *model.Submission) *gorm.DB {
	return db.Save(&submission)
}

func DeleteSubmission(db *gorm.DB, submissionID uint32) *gorm.DB {
	return db.Where("id = ?", submissionID).Delete(&model.Submission{})
}

// Contest
func CreateContest(db *gorm.DB, contest *model.Contest) *gorm.DB {
	return db.Create(&contest)
}

func GetContestByID(db *gorm.DB, id uint32) *gorm.DB {
	return db.Where("id = ?", id)
}

func GetContests(db *gorm.DB, page, pageSize int) *gorm.DB {
	return db.Offset((page - 1) * pageSize).Limit(pageSize)
}

func UpdateContest(db *gorm.DB, contest *model.Contest) *gorm.DB {
	return db.Save(&contest)
}

func DeleteContest(db *gorm.DB, contestID uint32) *gorm.DB {
	return db.Where("id = ?", contestID).Delete(&model.Contest{})
}

// Problem
func CreateProblem(db *gorm.DB, problem *model.Problem) *gorm.DB {
	return db.Create(&problem)
}

func GetProblemByID(db *gorm.DB, id uint32) *gorm.DB {
	return db.Where("id = ?", id)
}

func GetProblems(db *gorm.DB, page, pageSize int) *gorm.DB {
	return db.Offset((page - 1) * pageSize).Limit(pageSize)
}

func UpdateProblem(db *gorm.DB, problem *model.Problem) *gorm.DB {
	return db.Save(&problem)
}

func DeleteProblem(db *gorm.DB, problemID int) *gorm.DB {
	return db.Where("ProblemID = ?", problemID).Delete(&model.Problem{})
}

// User
func GetUsers(db *gorm.DB) *gorm.DB {
	return db
}

func GetUserByEmail(db *gorm.DB, email string) *gorm.DB {
	return db.Where("email = ?", email)
}

func GetUserByUserID(db *gorm.DB, userID uint32) *gorm.DB {
	return db.Where("id = ?", userID)
}

func GetUserByUsername(db *gorm.DB, username string) *gorm.DB {
	return db.Where("username = ?", username)
}

func CreateUser(db *gorm.DB, user *model.User) *gorm.DB {
	hashedPassword := lib.HashedPassword(user.Password)
	user.Password = string(hashedPassword)
	return db.Create(&user)
}

func UpdateUser(db *gorm.DB, user *model.User) *gorm.DB {
	return db.Save(&user)
}

func DeleteUser(db *gorm.DB, username string) *gorm.DB {
	return db.Where("username = ?", username).Delete(&model.User{})
}
