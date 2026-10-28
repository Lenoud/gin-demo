package main

import "github.com/gin-gonic/gin"

func main() {
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.String(200, "hello world!")
	})
	g.GET("/ping", func(c *gin.Context) {
		c.String(200, "online")
	})
	g.GET("/index", func(c *gin.Context) {
		c.String(200, "index page")
	})
	g.POST("/index", indexPost)
	if err := g.Run(":38080"); err != nil {
		panic(err)
	}
}

func indexPost(c *gin.Context) {
	c.String(200, "this is post method for gin")
}
