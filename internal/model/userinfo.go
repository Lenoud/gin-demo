package model

import (
	"time"

	"gorm.io/gorm"
)

type UserJson struct {
	Id        uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"username" binding:"required,min=3,max=20"`
	Password  string         `gorm:"type:varchar(255);not null" json:"password" binding:"required,min=6" json:"-"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex" json:"email" binding:"required,email"`
	IsAdmin   bool           `gorm:"default:false" json:"is_admin"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (UserJson) TableName() string {
	return "user_info"
}

func CreateUser(user *UserJson) error {
	return DB.Self.Create(user).Error
}

func GetUserByUsername(username string) (*UserJson, error) {
	var user UserJson
	err := DB.Self.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// 新增：列出所有用户
func ListUsers() ([]UserJson, error) {
	var users []UserJson
	err := DB.Self.Find(&users).Error
	return users, err
}
