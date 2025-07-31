package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 是程序的配置结构体
type Config struct {
	Name string // 配置文件路径（可选）
}

// InitConfig 初始化配置
func (c *Config) InitConfig() error {
	if c.Name != "" {
		// 如果指定了文件路径，直接加载
		viper.SetConfigFile(c.Name)
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

	fmt.Println("使用的配置文件:", viper.ConfigFileUsed())
	return nil
}
