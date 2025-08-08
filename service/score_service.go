package service

import (
	"errors"
	"time"

	"github.com/Lenoud/gin-demo/model"
)

func AddScore(studentID uint64, subject string, scoreValue float64) (*model.Score, error) {
	score := &model.Score{
		StudentId:  studentID,
		Subject:    subject,
		ScoreValue: scoreValue,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := model.CreateScore(score); err != nil {
		return nil, err
	}
	return score, nil
}

func GetScores(studentID uint64) ([]model.Score, error) {
	return model.GetScoresByStudentID(studentID)
}

func UpdateScore(scoreID uint64, subject string, scoreValue float64) (*model.Score, error) {
	updateData := map[string]interface{}{
		"updated_at": time.Now(),
	}
	if subject != "" {
		updateData["subject"] = subject
	}
	if scoreValue != 0 {
		updateData["score_value"] = scoreValue
	}

	if err := model.UpdateScoreByID(scoreID, updateData); err != nil {
		return nil, err
	}

	return model.GetScoreByID(scoreID)
}

func DeleteScore(scoreID uint64) error {
	exists, err := model.ScoreExists(scoreID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("成绩记录不存在")
	}
	return model.DeleteScoreByID(scoreID)
}
