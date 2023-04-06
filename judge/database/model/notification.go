package model

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID      uint32 `json:"id" gorm:"primaryKey;autoIncrement"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
