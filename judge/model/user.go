package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// don't mask the json field name, or it will lead to the shouldBindJSON ignore the field
	Username     string       `json:"username" gorm:"unique"`
	Password     string       `json:"password"`
	Email        string       `json:"email" gorm:"unique"`
	PersonalInfo PersonalInfo `json:"personal_info" gorm:"foreignKey:UserID"`
	AccountInfo  AccountInfo  `json:"account_info" gorm:"foreignKey:UserID"`
	Submissions  []Submission `json:"submissions" gorm:"foreignKey:UserID"`
	Ranks        []Rank       `json:"ranks" gorm:"foreignKey:UserID"`
}

type PersonalInfo struct {
	UserID       uint16 `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Birthday     string `json:"birthday"`
	PhoneNumber  string `json:"phone_number"`
	Introduction string `json:"introduction"`
}

type AccountInfo struct {
	UserID     uint16   `json:"user_id"`
	Permission uint8    `json:"permission"`
	IsActive   bool     `json:"is_active"`
	Tags       []string `json:"tags" gorm:"serializer:json"` // !!Serialize tags to json, or it will throw an error
}

type Rank struct {
	gorm.Model
	UserID     uint16 `json:"user_id"`
	ContestID  uint32 `json:"contest_id"`
	Rank       uint64 `json:"rank"`
	GroupID    uint32 `json:"group_id"`
	ClassRank  uint64 `json:"class_rank"`
	TotalUsers uint64 `json:"total_users"`
}
