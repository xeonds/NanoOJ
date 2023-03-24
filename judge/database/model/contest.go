package model

import (
	"time"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	ContestID  uint32
	ProblemIDs []uint32 `gorm:"many2many:contest_problems;"`
	Note       string
	Release    time.Time
	StartTime  time.Time
	EndTime    time.Time
}
