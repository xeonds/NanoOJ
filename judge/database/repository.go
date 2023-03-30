package database

import (
	"golang.org/x/crypto/bcrypt"
	"xyz.xeonds/nano-oj/database/model"
)

// Submissions
func CreateSubmission(submission *model.Submission) error {
	result := NanoDB.Create(&submission)
	return result.Error
}

func GetSubmissionByID(id uint32) (*model.Submission, error) {
	submission := &model.Submission{ID: id}
	err := NanoDB.First(&submission).Error
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func GetSubmissionsByProblemID(problemID uint32) ([]model.Submission, error) {
	var submissions []model.Submission
	err := NanoDB.Where("id = ?", problemID).Find(&submissions).Error
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func GetSubmissionsByUserID(userID uint32) ([]model.Submission, error) {
	var submissions []model.Submission
	err := NanoDB.Where("id = ?", userID).Find(&submissions).Error
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func GetSubmissionsByStatus(status string) ([]model.Submission, error) {
	var submissions []model.Submission
	err := NanoDB.Where("status = ?", status).Find(&submissions).Error
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func UpdateSubmission(submission *model.Submission) error {
	result := NanoDB.Save(&submission)
	return result.Error
}

func DeleteSubmission(submissionID uint32) error {
	result := NanoDB.Where("id = ?", submissionID).Delete(&model.Submission{})
	return result.Error
}

// Contest
func CreateContest(contest *model.Contest) error {
	result := NanoDB.Create(&contest)
	return result.Error
}

func GetContestByID(id uint32) (*model.Contest, error) {
	contest := &model.Contest{ID: id}
	err := NanoDB.First(&contest).Error
	if err != nil {
		return nil, err
	}
	return contest, nil
}

func GetContests(page, pageSize int) ([]*model.Contest, error) {
	contests := []*model.Contest{}
	offset := (page - 1) * pageSize
	err := NanoDB.Offset(offset).Limit(pageSize).Find(&contests).Error
	if err != nil {
		return nil, err
	}
	return contests, nil
}

func UpdateContest(contest *model.Contest) error {
	result := NanoDB.Save(&contest)
	return result.Error
}

func DeleteContest(contestID uint32) error {
	result := NanoDB.Where("id = ?", contestID).Delete(&model.Contest{})
	return result.Error
}

// Problem
func CreateProblem(problem *model.Problem) error {
	result := NanoDB.Create(&problem)
	return result.Error
}

func GetProblemByID(id uint32) (*model.Problem, error) {
	problem := &model.Problem{ID: id}
	err := NanoDB.First(&problem).Error
	if err != nil {
		return nil, err
	}
	return problem, nil
}

func GetProblems(page, pageSize int) ([]*model.Problem, error) {
	problems := []*model.Problem{}
	offset := (page - 1) * pageSize
	err := NanoDB.Offset(offset).Limit(pageSize).Find(&problems).Error
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func UpdateProblem(problem *model.Problem) error {
	result := NanoDB.Save(&problem)
	return result.Error
}

func DeleteProblem(problemID int) error {
	result := NanoDB.Where("ProblemID = ?", problemID).Delete(&model.Problem{})
	return result.Error
}

// User
func GetUsers() ([]model.User, error) {
	var users []model.User
	result := NanoDB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := NanoDB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByUserID(userid uint32) (*model.User, error) {
	var user model.User
	result := NanoDB.Where("id = ?", userid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := NanoDB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	result := NanoDB.Create(&user)
	return result.Error
}

func UpdateUser(user *model.User) error {
	result := NanoDB.Save(&user)
	return result.Error
}

func DeleteUser(username string) error {
	result := NanoDB.Where("username = ?", username).Delete(&model.User{})
	return result.Error
}
