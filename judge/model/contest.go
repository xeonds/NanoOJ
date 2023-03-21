package model

import (
	"time"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	Problems  []Problem `gorm:"many2many:contest_problems;"`
	Note      string
	Release   time.Time
	StartTime time.Time
	EndTime   time.Time
}
