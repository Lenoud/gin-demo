package user

import (
	"net/http"
	"strconv"

	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 绑定请求结构体
type BindRequest struct {
	UserId    uint64 `json:"user_id"`
	StudentId uint64 `json:"student_id"`
}

// 解绑请求结构体
type UnbindRequest struct {
	UserId uint64 `json:"user_id"`
}

// 绑定用户和学生接口
func BindUserStudent(c *gin.Context) {
	var req BindRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "请求参数错误", nil)
		return
	}
	if req.UserId == 0 || req.StudentId == 0 {
		controller.SendResponse(c, http.StatusBadRequest, "用户ID和学生ID不能为空", nil)
		return
	}

	if err := service.BindUserStudent(req.UserId, req.StudentId); err != nil {
		controller.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	controller.SendResponse(c, http.StatusOK, "绑定成功", gin.H{"bind": req})
}

// 解绑用户和学生接口
func UnbindUserStudent(c *gin.Context) {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil || userId == 0 {
		controller.SendResponse(c, http.StatusBadRequest, "无效的用户ID", nil)
		return
	}

	if err := service.UnbindUserStudent(userId); err != nil {
		controller.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	controller.SendResponse(c, http.StatusOK, "解绑成功", gin.H{"user_id": userId})
}

// ListUsers 获取所有用户列表接口
func ListUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		zap.L().Error("获取用户列表失败", zap.Error(err))
		controller.SendResponse(c, http.StatusInternalServerError, "获取用户列表失败", nil)
		return
	}
	controller.SendResponse(c, http.StatusOK, "成功", gin.H{"users": users})
}
