package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	Title         string    `json:"title"`
	Problems      []Problem `json:"problems" gorm:"many2many:contest_problems"`
	Description   string    `json:"description"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	IsRankVisible bool      `json:"is_rank_visible"`
}

type Notification struct {
	gorm.Model
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type Problem struct {
	gorm.Model
	Title       string  `json:"title"`
	Difficulty  int     `json:"difficulty"`
	Description string  `json:"description"`
	Inputs      Inputs  `json:"inputs"`
	Outputs     Outputs `json:"outputs"`
	Ranks       Ranks   `json:"ranks"`
	TimeLimit   int     `json:"time_limit"`
}

type Inputs []string

func (i *Inputs) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, i)
}
func (i Inputs) Value() (driver.Value, error) {
	return json.Marshal(i)
}

type Outputs []string

func (o *Outputs) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, o)
}
func (o Outputs) Value() (driver.Value, error) {
	return json.Marshal(o)
}

type Ranks []int

func (r *Ranks) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, r)
}
func (r Ranks) Value() (driver.Value, error) {
	return json.Marshal(r)
}

type Submission struct {
	gorm.Model
	ProblemID   uint32      `json:"problem_id"`
	Problem     Problem     `json:"problem"`
	UserID      uint16      `json:"user_id"`
	ContestID   uint        `json:"contest_id"`
	Language    string      `json:"language"`
	Code        string      `json:"code"`
	Status      Status      `json:"status"`
	Information Information `json:"information" gorm:"type:json"`
	Time        int         `json:"time"`
	Memory      int         `json:"memory"`
	Rank        int         `json:"rank"`
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
	UserID    uint16 `json:"user_id"`
	ContestID uint32 `json:"contest_id"` // means the user's rank of a problem in a contest
	ProblemID uint32 `json:"problem_id"`
	Rank      uint64 `json:"rank"`
}
