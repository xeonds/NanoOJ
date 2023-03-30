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
	Status      Status      `json:"status" gorm:"type:int"`
	Information Information `json:"information" gorm:"type:json"`
	Time        int         `json:"time"`
	Memory      int         `json:"memory"`
}

type Status int

const (
	Pending Status = iota
	InProgress
	Accepted
	WrongAnswer
	TimeLimitExceeded
	MemoryLimitExceeded
	RuntimeError
	CompilationError
)

func (s Status) String() string {
	switch s {
	case Pending:
		return "Pending"
	case InProgress:
		return "In progress"
	case Accepted:
		return "Accepted"
	case WrongAnswer:
		return "Wrong Answer"
	case TimeLimitExceeded:
		return "Time Limit Exceeded"
	case MemoryLimitExceeded:
		return "Memory Limit Exceeded"
	case RuntimeError:
		return "Runtime Error"
	case CompilationError:
		return "Compilation Error"
	default:
		return "Unknown"
	}
}

type Information []string

func (i *Information) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, i)
}

func (i Information) Value() (driver.Value, error) {
	return json.Marshal(i)
}
