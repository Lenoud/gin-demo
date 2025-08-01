package router


// URI 设计原则 
/*
使用名词而非动词表示资源
好：/users
不好：/getUsers
使用小写字母和连字符（-）
避免文件扩展名
使用复数形式表示集合
分层次表示关系：/users/{id}/orders
*/

import (
	"github.com/gin-gonic/gin"
	"github.com/Lenoud/gin-demo/controller/student"
	"github.com/Lenoud/gin-demo/controller/register"
)

func Load(g *gin.Engine){
	g.GET("/admin",func(c *gin.Context){
		c.String(200,"admin page!")
	})
	g.GET("/users",func(c *gin.Context){
		c.String(200,"user page!")
	})

	g.GET("/students", student.GetStudent)
	g.POST("/register",register.Register)
}



