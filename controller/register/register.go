package register

import (
	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.UserJson
	if err := c.ShouldBindJSON(&user); err != nil {
		controller.SendResponse(c, 400, "接收数据失败: "+err.Error(), nil)
		return
	}

	// 强制普通注册用户不是管理员
	// user.IsAdmin = false

	// 检查用户名是否存在
	if model.DB.Self.Where("username = ?", user.Username).First(&model.UserJson{}).RowsAffected > 0 {
		controller.SendResponse(c, 400, "用户名已存在", nil)
		return
	}

	if err := user.Create(); err != nil {
		controller.SendResponse(c, 500, "写入数据失败！", nil)
		return
	}
	controller.SendResponse(c, 200, "注册成功!", gin.H{"register": user})
}
