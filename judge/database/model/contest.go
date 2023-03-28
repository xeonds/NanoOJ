package model

import (
	"time"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	ContestID  uint32
	ProblemIDs []uint32 `gorm:"type:uint32"`
	Note       string
	Release    time.Time
	StartTime  time.Time
	EndTime    time.Time
}
