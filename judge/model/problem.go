package model

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	ID          uint32  `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string  `json:"title"`
	Difficulty  int     `json:"difficulty"`
	Description string  `json:"description"`
	Inputs      Inputs  `json:"inputs"`
	Outputs     Outputs `json:"outputs"`
	TimeLimit   int     `json:"time_limit"`
}

type Inputs []string
type Outputs []string

func (i *Inputs) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, i)
}

func (i Inputs) Value() (driver.Value, error) {
	return json.Marshal(i)
}

func (o *Outputs) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, o)
}

func (o Outputs) Value() (driver.Value, error) {
	return json.Marshal(o)
}
