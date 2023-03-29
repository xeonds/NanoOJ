package model

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	ID          uint32 `gorm:"primaryKey;autoIncrement"`
	Title       string
	Difficulty  int
	Description string
	Inputs      []string `gorm:"type:text"`
	Outputs     []string `gorm:"type:text"`
}
