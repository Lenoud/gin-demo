package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func OpenDB() *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.db"),
		true,
		"Local",
	)
	//config --> gin:123456@tcp(localhost)/gin_web?charset=utf8&parseTime=true&loc=Local
	db, err := gorm.Open("mysql", config)
	if err != nil {
		fmt.Println("Database connection faild.", err)
	}
	fmt.Println("Mysql连接成功！")
	return db
}

type Database struct {
	Self *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self: OpenDB(),
	}
}

// 关闭数据库
func (db *Database) Close() {
	DB.Self.Close()
}
