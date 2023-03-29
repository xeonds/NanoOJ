package model

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID      uint32 `gorm:"primaryKey;autoIncrement"`
	Title   string
	Author  string
	Content string
}
