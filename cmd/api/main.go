package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/Lenoud/gin-demo/internal/config"
	"github.com/Lenoud/gin-demo/internal/middleware"
	"github.com/Lenoud/gin-demo/internal/model"
	"github.com/Lenoud/gin-demo/internal/router"
	"github.com/Lenoud/gin-demo/internal/utils/logger"
)

func main() {

	// 读取配置文件
	conf := config.Config{ConfigName: "configs/config.yaml"}
	if err := conf.InitConfig(); err != nil {
		panic(err)
	}

	// 日志工具初始化
	logger.InitLogger()
	defer logger.Logger.Sync()

	// 创建engine
	g := gin.Default()

	//连接数据库
	model.DB = &model.Database{}
	if err := model.DB.Init(); err != nil {
		panic(err)
	}
	defer model.DB.Close()

	// 加载跨域中间件
	g.Use(middleware.CORSMiddleware())

	//加载路由
	router.Load(g)

	// 启动server
	g.Run(viper.GetString("server.ip") + ":" + viper.GetString("server.port"))
}
