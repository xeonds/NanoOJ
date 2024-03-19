package model

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type Submission struct {
	gorm.Model
	ID          uint32      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProblemID   uint32      `json:"problem_id"`
	Problem     Problem     `json:"problem"`
	UserID      uint16      `json:"user_id"`
	Language    string      `json:"language"`
	Code        string      `json:"code"`
	Status      Status      `json:"status"`
	Information Information `json:"information" gorm:"type:json"`
	Time        int         `json:"time"`
	Memory      int         `json:"memory"`
}

type Information []string

func (i *Information) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, i)
}

func (i Information) Value() (driver.Value, error) {
	return json.Marshal(i)
}

type Status string

const (
	Pending             Status = "Pending"
	InProgress          Status = "In Progress"
	Accepted            Status = "Accepted"
	WrongAnswer         Status = "Wrong Answer"
	TimeLimitExceeded   Status = "Time Limit Exceeded"
	MemoryLimitExceeded Status = "Memory Limit Exceeded"
	RuntimeError        Status = "Runtime Error"
	CompilationError    Status = "Compilation Error"
	Unknown             Status = "Unknown"
)

func (s Status) String() string {
	return string(s)
}
