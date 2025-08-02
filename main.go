package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/Lenoud/gin-demo/config"
	"github.com/Lenoud/gin-demo/model"
	"github.com/Lenoud/gin-demo/router"
)

func main() {
	g := gin.Default()

	conf := config.Config{"config/config.yaml"}
	if err := conf.InitConfig(); err != nil {
		panic(err)
	}

	//连接数据库
	model.DB.Init()
	model.DB.Self.LogMode(true)
	defer model.DB.Close()
	//加载路由
	router.Load(g)

	g.Run(viper.GetString("server.port"))
}
