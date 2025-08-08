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
	"github.com/Lenoud/gin-demo/controller/login"
	"github.com/Lenoud/gin-demo/controller/register"
	"github.com/Lenoud/gin-demo/controller/score"
	"github.com/Lenoud/gin-demo/controller/student"
	"github.com/Lenoud/gin-demo/controller/user"
	"github.com/Lenoud/gin-demo/middleware"
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) {

	g.POST("/api/auth/login", login.Login)

	api := g.Group("/api")
	//启用jwt鉴权
	api.Use(middleware.JWTAuthMiddleware())
	{
		//学生相关接口
		api.GET("/students", student.GetStudents)
		api.POST("/students", student.AddStudent)
		api.DELETE("/students/:id", student.DelStudent)
		api.PUT("/students/:id", student.UpdateStudent)

		// 学生成绩相关
		api.POST("/students/:id/scores", score.AddScore) // 添加成绩
		api.GET("/students/:id/scores", score.GetScores) // 查询成绩
		api.PUT("/scores/:score_id", score.UpdateScore)  // 修改成绩
		api.DELETE("/scores/:score_id", score.DelScore)  // 删除成绩

		// 用户注册
		api.POST("/register", register.Register)

		// 用户的管理
		api.GET("/users", user.ListUsers)
		// api.GET("/users/:id", register.Register)
		// api.DELETE("/users/:id", register.Register)
		// api.POST("/users/:id/students", register.Register)
		// 绑定用户和学生接口（需要鉴权）
		api.POST("/user_students/bind", user.BindUserStudent)
		api.DELETE("/user_students/bind/:user_id", user.UnbindUserStudent)

	}
}
