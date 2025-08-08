package login

import (
	"net/http"

	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/model"
	"github.com/Lenoud/gin-demo/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		controller.SendResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error(), nil)
		return
	}

	token, user, err := service.LoginUser(req.Username, req.Password)
	if err != nil {
		controller.SendResponse(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

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
