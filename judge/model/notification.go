package model

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
