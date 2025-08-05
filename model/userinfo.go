package model

import (
	"time"
)

// UserJson 结构体，存储用户信息
type UserJson struct {
	Id        uint64    `gorm:"column:id;primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"column:username;type:varchar(50);uniqueIndex;not null" json:"username" binding:"required,min=3,max=20"`
	Password  string    `gorm:"column:password;type:varchar(255);not null" json:"password" binding:"required,min=6" json:"-"`
	Email     string    `gorm:"column:email;type:varchar(100);uniqueIndex" json:"email" binding:"required,email"`
	IsAdmin   bool      `gorm:"column:is_admin;default:false" json:"is_admin"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (u *UserJson) TableName() string {
	return "user_info"
}

func (u *UserJson) Create() error {
	return DB.Self.Create(&u).Error
}
