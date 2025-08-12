package register

import (
	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/model"
	"github.com/Lenoud/gin-demo/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Register(c *gin.Context) {
	zap.L().Info("收到用户注册请求")
	var user model.UserJson
	if err := c.ShouldBindJSON(&user); err != nil {
		controller.SendResponse(c, 400, "接收数据失败: "+err.Error(), nil)
		zap.L().Warn("接收用户注册数据失败", zap.String("error", err.Error()))
		return
	}

	newUser, err := service.RegisterUser(&user)
	if err != nil {
		controller.SendResponse(c, 400, "注册失败: "+err.Error(), nil)
		zap.L().Error("用户注册失败", zap.String("username", user.Username), zap.String("error", err.Error()))
		return
	}

	controller.SendResponse(c, 200, "注册成功!", gin.H{"register": newUser})
}
