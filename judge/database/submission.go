package database

import "xyz.xeonds/nano-oj/model"

func GetSubmissionsByStatus(status string) ([]model.Submission, error) {
	var submissions []model.Submission
	err := NanoDB.Select(&submissions, "SELECT * FROM submissions WHERE status = ?", status).Error
	if err != nil {
		return nil, err
	}
	return submissions, nil
}
