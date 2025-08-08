package score

import (
	"net/http"
	"strconv"

	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/service"
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

	score, err := service.AddScore(studentID, req.Subject, req.ScoreValue)
	if err != nil {
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

	scores, err := service.GetScores(studentID)
	if err != nil {
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

	score, err := service.UpdateScore(scoreID, req.Subject, req.ScoreValue)
	if err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "修改成绩失败", nil)
		return
	}

	controller.SendResponse(c, http.StatusOK, "成绩修改成功", score)
}

// DelScore 删除成绩
func DelScore(c *gin.Context) {
	scoreIDStr := c.Param("score_id")
	scoreID, err := strconv.ParseUint(scoreIDStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "成绩ID格式错误", nil)
		return
	}

	if err := service.DeleteScore(scoreID); err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "删除成绩失败", nil)
		return
	}

	controller.SendResponse(c, http.StatusOK, "成绩删除成功", scoreID)
}
