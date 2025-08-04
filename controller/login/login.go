package login

import (
	"net/http"

	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/model"
	"github.com/Lenoud/gin-demo/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error(), nil)
		return
	}

	var user model.UserJson
	err := model.DB.Self.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		// 用户不存在
		controller.SendResponse(c, http.StatusUnauthorized, "用户名或密码错误", nil)
		return
	}

	// 明文密码比较
	if user.Password != req.Password {
		controller.SendResponse(c, http.StatusUnauthorized, "用户名或密码错误", nil)
		return
	}

	// 生成 JWT Token
	token, err := utils.GenerateToken(user.Id, user.IsAdmin)
	if err != nil {
		controller.SendResponse(c, http.StatusInternalServerError, "生成Token失败", nil)
		return
	}

	// 登录成功，返回用户信息
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
