package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   uint16 `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"uniqueIndex"`
	Password string
	Email    string `gorm:"unique"`
}
