package database

import (
	"xyz.xeonds/nano-oj/model"
)

func CreateProblem(problem *model.Problem) error {
	result := NanoDB.Create(&problem)
	return result.Error
}

func GetProblemByID(id uint32) (*model.Problem, error) {
	problem := &model.Problem{ProblemID: id}
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
