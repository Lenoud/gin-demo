package model

import "time"

type Score struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentId  uint64    `gorm:"not null" json:"student_id"`
	Subject    string    `gorm:"not null" json:"subject"`
	ScoreValue float64   `gorm:"not null" json:"score_value"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Score) TableName() string {
	return "scores"
}
