package service

import (
	"errors"
	"time"

	"github.com/Lenoud/gin-demo/model"
	"go.uber.org/zap"
)

func AddScore(studentID uint64, subject string, scoreValue float64) (*model.Score, error) {
	zap.L().Info("尝试添加成绩", zap.Uint64("studentId", studentID), zap.String("subject", subject), zap.Float64("scoreValue", scoreValue))
	score := &model.Score{
		StudentId:  studentID,
		Subject:    subject,
		ScoreValue: scoreValue,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := model.CreateScore(score); err != nil {
		zap.L().Error("添加成绩失败", zap.Uint64("studentId", studentID), zap.String("subject", subject), zap.String("error", err.Error()))
		return nil, err
	}
	zap.L().Info("成绩添加成功", zap.Uint64("studentId", studentID), zap.String("subject", subject), zap.Float64("scoreValue", scoreValue))
	return score, nil
}

func GetScores(studentID uint64) ([]model.Score, error) {
	zap.L().Info("尝试获取学生成绩", zap.Uint64("studentId", studentID))
	scores, err := model.GetScoresByStudentID(studentID)
	if err != nil {
		zap.L().Error("获取学生成绩失败", zap.Uint64("studentId", studentID), zap.String("error", err.Error()))
		return nil, err
	}
	return scores, nil
}

func UpdateScore(scoreID uint64, subject string, scoreValue float64) (*model.Score, error) {
	zap.L().Info("尝试更新成绩", zap.Uint64("scoreId", scoreID), zap.String("subject", subject), zap.Float64("scoreValue", scoreValue))
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
		zap.L().Error("更新成绩失败", zap.Uint64("scoreId", scoreID), zap.String("error", err.Error()))
		return nil, err
	}

	zap.L().Info("成绩更新成功", zap.Uint64("scoreId", scoreID), zap.String("subject", subject), zap.Float64("scoreValue", scoreValue))
	return model.GetScoreByID(scoreID)
}

func DeleteScore(scoreID uint64) error {
	zap.L().Info("尝试删除成绩", zap.Uint64("scoreId", scoreID))
	exists, err := model.ScoreExists(scoreID)
	if err != nil {
		zap.L().Error("删除成绩失败", zap.Uint64("scoreId", scoreID), zap.String("error", err.Error()))
		return err
	}
	if !exists {
		return errors.New("成绩记录不存在")
	}
	if err := model.DeleteScoreByID(scoreID); err != nil {
		zap.L().Error("删除成绩失败", zap.Uint64("scoreId", scoreID), zap.String("error", err.Error()))
		return err
	}
	zap.L().Info("成绩删除成功", zap.Uint64("scoreId", scoreID))
	return nil
}
