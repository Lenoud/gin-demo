package score

import (
	"net/http"
	"strconv"

	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AddScore 给某个学生添加成绩
func AddScore(c *gin.Context) {
	zap.L().Info("收到添加成绩请求", zap.String("studentId", c.Param("id")))
	studentIDStr := c.Param("id")
	studentID, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "学生ID格式错误", nil)
		zap.L().Warn("学生ID格式错误", zap.String("studentId", studentIDStr), zap.String("error", err.Error()))
		return
	}

	var req struct {
		Subject    string  `json:"subject" binding:"required"`
		ScoreValue float64 `json:"score_value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "请求参数错误", nil)
		zap.L().Warn("请求参数错误", zap.String("error", err.Error()))
		return
	}

	score, err := service.AddScore(studentID, req.Subject, req.ScoreValue)
	if err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "添加成绩失败", nil)
		zap.L().Error("添加成绩失败", zap.Uint64("studentId", studentID), zap.String("error", err.Error()))
		return
	}

	controller.SendResponse(c, http.StatusOK, "成绩添加成功", score)
	zap.L().Info("成绩添加成功", zap.Uint64("studentId", studentID), zap.String("subject", req.Subject), zap.Float64("scoreValue", req.ScoreValue))
}

// GetScores 查询某个学生的成绩
func GetScores(c *gin.Context) {
	zap.L().Info("收到获取成绩请求", zap.String("studentId", c.Param("id")))
	studentIDStr := c.Param("id")
	studentID, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "学生ID格式错误", nil)
		zap.L().Warn("学生ID格式错误", zap.String("studentId", studentIDStr), zap.String("error", err.Error()))
		return
	}

	scores, err := service.GetScores(studentID)
	if err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "查询成绩失败", nil)
		zap.L().Error("查询成绩失败", zap.Uint64("studentId", studentID), zap.String("error", err.Error()))
		return
	}

	controller.SendResponse(c, http.StatusOK, "查询成功", scores)
	zap.L().Info("成绩查询成功", zap.Uint64("studentId", studentID), zap.Int("scoreCount", len(scores)))
}

// UpdateScore 修改成绩
func UpdateScore(c *gin.Context) {
	zap.L().Info("收到更新成绩请求", zap.String("scoreId", c.Param("score_id")))
	scoreIDStr := c.Param("score_id")
	scoreID, err := strconv.ParseUint(scoreIDStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "成绩ID格式错误", nil)
		zap.L().Warn("成绩ID格式错误", zap.String("scoreId", scoreIDStr), zap.String("error", err.Error()))
		return
	}

	var req struct {
		Subject    string  `json:"subject"`
		ScoreValue float64 `json:"score_value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "请求参数错误", nil)
		zap.L().Warn("请求参数错误", zap.String("error", err.Error()))
		return
	}

	score, err := service.UpdateScore(scoreID, req.Subject, req.ScoreValue)
	if err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "修改成绩失败", nil)
		zap.L().Error("修改成绩失败", zap.Uint64("scoreId", scoreID), zap.String("error", err.Error()))
		return
	}

	controller.SendResponse(c, http.StatusOK, "成绩修改成功", score)
	zap.L().Info("成绩更新成功", zap.Uint64("scoreId", scoreID), zap.String("subject", req.Subject), zap.Float64("scoreValue", req.ScoreValue))
}

// DelScore 删除成绩
func DelScore(c *gin.Context) {
	zap.L().Info("收到删除成绩请求", zap.String("scoreId", c.Param("score_id")))
	scoreIDStr := c.Param("score_id")
	scoreID, err := strconv.ParseUint(scoreIDStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "成绩ID格式错误", nil)
		zap.L().Warn("成绩ID格式错误", zap.String("scoreId", scoreIDStr), zap.String("error", err.Error()))
		return
	}

	if err := service.DeleteScore(scoreID); err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "删除成绩失败", nil)
		zap.L().Error("删除成绩失败", zap.Uint64("scoreId", scoreID), zap.String("error", err.Error()))
		return
	}

	controller.SendResponse(c, http.StatusOK, "成绩删除成功", scoreID)
	zap.L().Info("成绩删除成功", zap.Uint64("scoreId", scoreID))
}
