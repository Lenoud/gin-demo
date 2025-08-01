package register

import (
	"github.com/gin-gonic/gin"
	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/model"
)

func Register(c *gin.Context){
	var user model.UserJson
	if err:= c.ShouldBindJSON(&user);err != nil {
		controller.SendResponse(c, 400, "注册失败: "+err.Error(), nil)
		return
	}
	controller.SendResponse(c,200,"注册成功!",gin.H{"register": user})
}