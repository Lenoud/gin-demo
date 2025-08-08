package register

import (
	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/model"
	"github.com/Lenoud/gin-demo/service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.UserJson
	if err := c.ShouldBindJSON(&user); err != nil {
		controller.SendResponse(c, 400, "接收数据失败: "+err.Error(), nil)
		return
	}

	newUser, err := service.RegisterUser(&user)
	if err != nil {
		controller.SendResponse(c, 400, "注册失败: "+err.Error(), nil)
		return
	}

	controller.SendResponse(c, 200, "注册成功!", gin.H{"register": newUser})
}
