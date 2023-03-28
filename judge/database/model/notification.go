package model

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID         uint32
	Title      string
	Author     string
	Content    string
	Release    time.Time
	UpdateTime time.Time
}
