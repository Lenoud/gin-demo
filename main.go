package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"


	"github.com/Lenoud/gin-demo/config"
	"github.com/Lenoud/gin-demo/router"
)

func main(){
	//创建engine
	g:= gin.Default()

	//配置静态文件和html文件位置
	g.Static("/static","./static")
	g.LoadHTMLGlob("templates/*")

	// 读取配置
	conf := config.Config{"config/config.yaml"}
	if err := conf.InitConfig();err!=nil{
		panic(err)
	}

	//加载路由
	router.Load(g)
	
	g.Run(viper.GetString("server.port"))



}