package model

import (
	"time"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	ID          uint32    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title"`
	Problems    []Problem `json:"problems" gorm:"many2many:contest_problems"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
