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
	"github.com/Lenoud/gin-demo/controller/register"
	"github.com/Lenoud/gin-demo/controller/student"
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) {

	// g.Use(cors.New(cors.Config{
	// 	// 允许的前端来源（必须精确匹配你的前端地址）
	// 	AllowOrigins: []string{"http://192.168.100.153:5173"},
	// 	// 允许的请求方法
	// 	AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	// 允许的请求头
	// 	AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	// 	// 允许暴露的响应头
	// 	ExposeHeaders: []string{"Content-Length"},
	// 	// 是否允许携带凭证（如 cookies）
	// 	AllowCredentials: true,
	// 	// 预检请求的缓存时间（减少 OPTIONS 请求）
	// 	MaxAge: 12 * time.Hour,
	// }))

	g.GET("/admin", func(c *gin.Context) {
		c.String(200, "admin page!")
	})
	g.GET("/users", func(c *gin.Context) {
		c.String(200, "user page!")
	})
	g.POST("/register", register.Register)

	api := g.Group("/api")
	{
		api.GET("/students", student.GetStudents)        // 接口路径变为 /api/students
		api.POST("/addstu", student.AddStudent)          // 接口路径变为 /api/addstu
		api.DELETE("/delstu/:id", student.DelStudent)    // 接口路径变为 /api/delstu/:id
		api.PUT("/updatestu/:id", student.UpdateStudent) // 接口路径变为 /api/updatestu/:id
	}

}
