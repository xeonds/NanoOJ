package model

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	ProblemID uint
	UserID    uint
	Language  string
	Code      string
	Status    string
	Time      int
	Memory    int
}
