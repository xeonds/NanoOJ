package database

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"xyz.xeonds/nano-oj/database/model"
)

type Repository struct {
	DB *gorm.DB
}

// Submissions
func (r *Repository) CreateSubmission(submission *model.Submission) error {
	result := r.DB.Create(&submission)
	return result.Error
}

func (r *Repository) GetSubmissionByID(id uint32) (*model.Submission, error) {
	submission := &model.Submission{ID: id}
	err := r.DB.First(&submission).Error
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (r *Repository) GetSubmissionsByProblemID(problemID uint32) ([]model.Submission, error) {
	var submissions []model.Submission
	err := r.DB.Where("id = ?", problemID).Find(&submissions).Error
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func (r *Repository) GetSubmissionsByUserID(userID uint32) ([]model.Submission, error) {
	var submissions []model.Submission
	err := r.DB.Where("id = ?", userID).Find(&submissions).Error
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func (r *Repository) GetSubmissionsByStatus(status model.Status) ([]model.Submission, error) {
	var submissions []model.Submission
	err := r.DB.Where("status = ?", status).Find(&submissions).Error
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func (r *Repository) UpdateSubmission(submission *model.Submission) error {
	result := r.DB.Save(&submission)
	return result.Error
}

func (r *Repository) DeleteSubmission(submissionID uint32) error {
	result := r.DB.Where("id = ?", submissionID).Delete(&model.Submission{})
	return result.Error
}

// Contest
func (r *Repository) CreateContest(contest *model.Contest) error {
	result := r.DB.Create(&contest)
	return result.Error
}

func (r *Repository) GetContestByID(id uint32) (*model.Contest, error) {
	contest := &model.Contest{ID: id}
	err := r.DB.First(&contest).Error
	if err != nil {
		return nil, err
	}
	return contest, nil
}

func (r *Repository) GetContests(page, pageSize int) ([]*model.Contest, error) {
	contests := []*model.Contest{}
	offset := (page - 1) * pageSize
	err := r.DB.Offset(offset).Limit(pageSize).Find(&contests).Error
	if err != nil {
		return nil, err
	}
	return contests, nil
}

func (r *Repository) UpdateContest(contest *model.Contest) error {
	result := r.DB.Save(&contest)
	return result.Error
}

func (r *Repository) DeleteContest(contestID uint32) error {
	result := r.DB.Where("id = ?", contestID).Delete(&model.Contest{})
	return result.Error
}

// Problem
func (r *Repository) CreateProblem(problem *model.Problem) error {
	result := r.DB.Create(&problem)
	return result.Error
}

func (r *Repository) GetProblemByID(id uint32) (*model.Problem, error) {
	problem := &model.Problem{ID: id}
	err := r.DB.First(&problem).Error
	if err != nil {
		return nil, err
	}
	return problem, nil
}

func (r *Repository) GetProblems(page, pageSize int) ([]*model.Problem, error) {
	problems := []*model.Problem{}
	offset := (page - 1) * pageSize
	err := r.DB.Offset(offset).Limit(pageSize).Find(&problems).Error
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func (r *Repository) UpdateProblem(problem *model.Problem) error {
	result := r.DB.Save(&problem)
	return result.Error
}

func (r *Repository) DeleteProblem(problemID int) error {
	result := r.DB.Where("ProblemID = ?", problemID).Delete(&model.Problem{})
	return result.Error
}

// User
func (r *Repository) GetUsers() ([]model.User, error) {
	var users []model.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *Repository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) GetUserByUserID(userid uint32) (*model.User, error) {
	var user model.User
	result := r.DB.Where("id = ?", userid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) CreateUser(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	result := r.DB.Create(&user)
	return result.Error
}

func (r *Repository) UpdateUser(user *model.User) error {
	result := r.DB.Save(&user)
	return result.Error
}

func (r *Repository) DeleteUser(username string) error {
	result := r.DB.Where("username = ?", username).Delete(&model.User{})
	return result.Error
}
