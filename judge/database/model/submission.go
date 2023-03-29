package model

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	SubmissionID uint32 `gorm:"primaryKey;autoIncrement"`
	ProblemID    uint32
	UserID       uint16
	Language     string
	Code         string
	Status       string
	Time         int
	Memory       int
}
