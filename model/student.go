package model

type StudentInfo struct {
	Id         uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	StuName    string `gorm:"not null" json:"stu_name"`
	StuAge     uint8  `gorm:"not null" json:"stu_age"`
	StuSex     string `gorm:"not null" json:"stu_sex"`
	StuAddress  string `gorm:"not null" json:"stu_address"`
}

func (s *StudentInfo) TableName() string {
	return "student_info"
}

func (s *StudentInfo) Create() error {
	return DB.Self.Create(&s).Error
}

// 在StudentInfo结构体中添加Delete方法
func (s *StudentInfo) Delete() error {
	return DB.Self.Delete(s).Error
}

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
