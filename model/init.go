package model

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.db"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 这里可根据需要设置日志级别
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %v", err)
	}
	fmt.Println("Mysql连接成功！数据库配置:", viper.GetStringMap("database"))
	return db, nil
}

type Database struct {
	Self *gorm.DB
}

var DB *Database

func (db *Database) Init() error {
	selfDB, err := OpenDB()
	if err != nil {
		return err
	}

	// 自动迁移
	err = selfDB.AutoMigrate(
		&UserJson{},
		&StudentInfo{},
		&UserStudent{},
		&Score{},
	)
	if err != nil {
		return fmt.Errorf("数据表创建失败: %v", err)
	}

	// 如果不存在 admin 用户，则创建一个（注意密码应加密）
	var count int64
	selfDB.Model(&UserJson{}).Where("username = ?", "admin").Count(&count)
	if count == 0 {
		adminUser := UserJson{
			Username: "admin",
			Password: "admin123", // 生产环境请使用加密密码
			IsAdmin:  true,
			Email:    "admin@example.com",
		}
		if err := selfDB.Create(&adminUser).Error; err != nil {
			return fmt.Errorf("创建管理员用户失败: %v", err)
		}
		fmt.Println("默认管理员 admin 用户已创建")
	}

	DB = &Database{
		Self: selfDB,
	}

	fmt.Println("数据库初始化完成，数据表已自动创建")
	return nil
}

func (db *Database) Close() error {
	sqlDB, err := DB.Self.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
