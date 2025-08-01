package router

import (
	"fmt"
	"net/http"

	"github.com/Lenoud/gin-demo/model"
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	g.POST("/login", func(c *gin.Context) {
		user := c.PostForm("username")
		password := c.PostForm("password")
		if user == "admin" && password == "123456" {
			c.String(http.StatusOK, "welcome to login page\nuser: %s, password: %s", user, password)
			fmt.Println(user, password)
		} else {
			// 渲染html模板数据
			c.HTML(http.StatusUnauthorized, "index.html", gin.H{"error": "用户名或密码错误"})
		}
	})

	//动态url
	g.POST("/login/:id", func(c *gin.Context) {
		id := c.Param("id") // 从 URL 中获取 id
		user := c.PostForm("username")
		password := c.PostForm("password")
		if user == "admin" && password == "123456" {
			c.String(http.StatusOK, "welcome to login page\nuser: %s, password: %s, id: %s", user, password, id)
			fmt.Println(user, password, id)
		} else {
			// 渲染html模板数据
			c.HTML(http.StatusUnauthorized, "index.html", gin.H{"error": "用户名或密码错误"})
		}
	})

	//注册路由
	g.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	g.POST("/register", register)

}

// 注册函数
func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{Username: username, Password: password}

	model.Users[username] = user
	fmt.Println(user, "\n", model.Users)
	c.String(http.StatusOK, "注册成功!\n用户名：%s,密码：%s", username, password)
	

	//接收json格式的数据
	var user1 model.UserJson

	if err := c.ShouldBindJSON(&user1); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
		return
	}
	fmt.Printf("%+v",user1)
	c.JSON(200,gin.H{"user":user1})

}
