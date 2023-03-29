package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         uint16 `gorm:"primaryKey;autoIncrement"`
	Username   string `gorm:"unique"`
	Permission uint8
	Password   string
	Email      string `gorm:"unique"`
}
