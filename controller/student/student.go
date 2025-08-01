package student

import (
	"github.com/Lenoud/gin-demo/controller"
	"github.com/gin-gonic/gin"
)


func GetStudent(c *gin.Context){
	controller.SendResponse(c,200,"成功!",nil)
	return
}
