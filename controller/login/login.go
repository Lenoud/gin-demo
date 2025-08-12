package login

import (
	"net/http"

	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
	zap.L().Info("收到用户登录请求")
	type LoginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error(), nil)
		zap.L().Warn("接收用户登录数据失败", zap.String("error", err.Error()))
		return
	}

	token, user, err := service.LoginUser(req.Username, req.Password)
	if err != nil {
		controller.SendResponse(c, http.StatusUnauthorized, err.Error(), nil)
		zap.L().Warn("用户登录失败", zap.String("username", req.Username), zap.String("error", err.Error()))
		return
	}

	zap.L().Info("用户登录成功", zap.String("username", req.Username), zap.Uint64("userId", user.Id))

	controller.SendResponse(c, http.StatusOK, "登录成功", gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.Id,
			"username": user.Username,
			"email":    user.Email,
			"is_admin": user.IsAdmin,
		},
	})
}
