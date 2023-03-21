package model

import "gorm.io/gorm"

type Problem struct {
    gorm.Model
	ProblemID      int    `gorm:"uniqueIndex"`
    ProblemNote    string
	ProblemInput   string
	ExpectedOutput string
}