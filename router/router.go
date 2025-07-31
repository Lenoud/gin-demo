package router

import (
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine){
	g.GET("/",func(c *gin.Context){
		c.String(200,"hello,i am gin!")
	})
	g.GET("/stu",func(c *gin.Context){
		c.String(200,"student!")
	})
}



