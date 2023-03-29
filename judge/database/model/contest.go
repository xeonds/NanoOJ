package model

import (
	"time"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	ID          uint32   `gorm:"primaryKey;autoIncrement"`
	ProblemIDs  []uint32 `gorm:"type:uint32"`
	Description string
	StartTime   time.Time
	EndTime     time.Time
}
