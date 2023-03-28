package model

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	ProblemID          uint32 `gorm:"uniqueIndex"`
	ProblemTitle       string
	ProblemDifficulty  int
	ProblemDescription string
	ProblemInputs      []string `gorm:"type:text"`
	ExpectedOutputs    []string `gorm:"type:text"`
}