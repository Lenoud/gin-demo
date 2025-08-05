package score

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/model"
	"github.com/gin-gonic/gin"
)

// AddScore 给某个学生添加成绩
func AddScore(c *gin.Context) {
	studentIDStr := c.Param("id")
	studentID, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "学生ID格式错误", nil)
		return
	}

	var req struct {
		Subject    string  `json:"subject" binding:"required"`
		ScoreValue float64 `json:"score_value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "请求参数错误", nil)
		return
	}

	score := model.Score{
		StudentId:  studentID,
		Subject:    req.Subject,
		ScoreValue: req.ScoreValue,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := model.DB.Self.Create(&score).Error; err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "添加成绩失败", nil)
		return
	}
	controller.SendResponse(c, http.StatusOK, "成绩添加成功", score)
}

// GetScores 查询某个学生的成绩
func GetScores(c *gin.Context) {
	studentIDStr := c.Param("id")
	studentID, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "学生ID格式错误", nil)
		return
	}

	var scores []model.Score
	if err := model.DB.Self.Where("student_id = ?", studentID).Find(&scores).Error; err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "查询成绩失败", nil)
		return
	}

	controller.SendResponse(c, http.StatusOK, "查询成功", scores)
}

// UpdateScore 修改成绩
func UpdateScore(c *gin.Context) {
	scoreIDStr := c.Param("score_id")
	scoreID, err := strconv.ParseUint(scoreIDStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "成绩ID格式错误", nil)
		return
	}

	var req struct {
		Subject    string  `json:"subject"`
		ScoreValue float64 `json:"score_value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "请求参数错误", nil)
		return
	}

	updateData := map[string]interface{}{
		"updated_at": time.Now(),
	}
	if req.Subject != "" {
		updateData["subject"] = req.Subject
	}
	if req.ScoreValue != 0 {
		updateData["score_value"] = req.ScoreValue
	}

	if err := model.DB.Self.Model(&model.Score{}).Where("id = ?", scoreID).Updates(updateData).Error; err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "修改成绩失败", nil)
		return
	}

	// 查询更新后的成绩信息并返回
	var updatedScore model.Score
	if err := model.DB.Self.First(&updatedScore, scoreID).Error; err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "获取更新后成绩失败", nil)
		return
	}
	
	controller.SendResponse(c, http.StatusOK, "成绩修改成功", updatedScore)
}

// DelScore 删除成绩
func DelScore(c *gin.Context) {
	scoreIDStr := c.Param("score_id")
	scoreID, err := strconv.ParseUint(scoreIDStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "成绩ID格式错误", nil)
		return
	}

	// 先查询是否存在
	var score model.Score
	if err := model.DB.Self.First(&score, scoreID).Error; err != nil {
		controller.SendResponse(c, http.StatusNotFound, "成绩记录不存在", nil)
		return
	}

	if err := model.DB.Self.Delete(&model.Score{}, scoreID).Error; err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "删除成绩失败", nil)
		return
	}

	controller.SendResponse(c, http.StatusOK, "成绩删除成功", scoreID)
}
