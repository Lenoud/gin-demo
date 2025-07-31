package main
import "github.com/gin-gonic/gin"

// func main(){
// 	g := gin.Default()
// 	g.GET("/",func(c *gin.Context){
// 		c.String(200,"hello world!")
// 	})
// 	g.GET("/ping",func(c *gin.Context){
// 		c.String(200,"online")
// 	})
// 	g.GET("/index",func(c *gin.Context){
// 		c.String(200,"index page")
// 	})
// 	g.POST("/index",index_post)
// 	g.Run(":38080")
// }

func index_post(c *gin.Context){
	c.String(200,"this is post method for gin")
}