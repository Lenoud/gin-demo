package model

import (
	"time"

	"gorm.io/gorm"
)

type Score struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentId  uint64    `gorm:"not null" json:"student_id"`
	Subject    string    `gorm:"not null;" json:"subject"`
	ScoreValue float64   `gorm:"not null" json:"score_value"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Score) TableName() string {
	return "scores"
}

func CreateScore(score *Score) error {
	return DB.Self.Create(score).Error
}

func GetScoresByStudentID(studentID uint64) ([]Score, error) {
	var scores []Score
	err := DB.Self.Where("student_id = ?", studentID).Find(&scores).Error
	return scores, err
}

func UpdateScoreByID(scoreID uint64, updateData map[string]interface{}) error {
	return DB.Self.Model(&Score{}).Where("id = ?", scoreID).Updates(updateData).Error
}

func GetScoreByID(scoreID uint64) (*Score, error) {
	var score Score
	if err := DB.Self.First(&score, scoreID).Error; err != nil {
		return nil, err
	}
	return &score, nil
}

func DeleteScoreByID(scoreID uint64) error {
	return DB.Self.Delete(&Score{}, scoreID).Error
}

func ScoreExists(scoreID uint64) (bool, error) {
	var score Score
	err := DB.Self.First(&score, scoreID).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}
