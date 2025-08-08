package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config 是程序的配置结构体
type Config struct {
	ConfigName string // 配置文件路径（可选）
}

// InitConfig 初始化配置
func (c *Config) InitConfig() error {
	if c.ConfigName != "" {
		// 如果指定了文件路径，直接加载
		viper.SetConfigFile(c.ConfigName)
	} else {
		// 默认加载 config/config.yaml
		viper.AddConfigPath("config") // 搜索路径
		viper.SetConfigName("config") // 文件名（不带扩展名）
		viper.SetConfigType("yaml")   // 文件类型
	}

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 环境变量覆盖机制 export SERVER_IP=127.0.0.1 可以覆盖掉server.ip 的值
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// 打印日志
	fmt.Println("使用的配置文件:", viper.ConfigFileUsed())
	return nil
}
