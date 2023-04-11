package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint16       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username    string       `json:"username" gorm:"unique"`
	Permission  uint8        `json:"permission"`
	Password    string       `json:"password"`
	Email       string       `json:"email" gorm:"unique"`
	Submissions []Submission `json:"submissions" gorm:"foreignKey:UserID"`
	Rank        uint64       `json:"rank"`
}
