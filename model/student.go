package model

import "errors"

type StudentInfo struct {
	Id         uint64 `gorm:"column:id;primary_key;auto_increment" json:"id"`
	StrName    string `gorm:"column:stu_name;not null" json:"stu_name"  binding:"required"`
	StrAge     uint8  `gorm:"column:stu_age;not null" json:"stu_age" binding:"required"`
	StuSex     string `gorm:"column:stu_sex;not null" json:"stu_sex" binding:"required"`
	StuAdderss string `gorm:"column:stu_address;not null" json:"stu_address" binding:"required"`
}

func (s *StudentInfo) TableName() string {
	return "student_info"
}

func (s *StudentInfo) Create() error {
	return DB.Self.Create(&s).Error
}

// 在StudentInfo结构体中添加Delete方法
func (s *StudentInfo) Delete() error {
	result := DB.Self.Delete(s)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

// 在model包中定义未找到记录的错误
var ErrRecordNotFound = errors.New("record not found")

// 获取所有学生信息
func (s *StudentInfo) GetAll() ([]StudentInfo, error) {
	var students []StudentInfo
	result := DB.Self.Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}

// 添加更新方法到StudentInfo结构体
func (s *StudentInfo) Update(updateData *StudentInfo) error {
	// 先查询记录是否存在
	if err := DB.Self.First(s, s.Id).Error; err != nil {
		return err
	}
	// 执行更新操作
	result := DB.Self.Model(s).Updates(updateData)
	return result.Error
}
