package model

type UserStudent struct {
	UserId    uint64 `gorm:"primaryKey;autoIncrement:false" json:"user_id"`
	StudentId uint64 `gorm:"not null" json:"student_id"`
}

func (UserStudent) TableName() string {
	return "user_students"
}

// 是否存在绑定关系
func ExistsUserStudent(userId, studentId uint64) (bool, error) {
	var count int64
	if err := DB.Self.Model(&UserStudent{}).
		Where("user_id = ? AND student_id = ?", userId, studentId).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// 创建绑定关系
func CreateUserStudent(userId, studentId uint64) error {
	bind := UserStudent{UserId: userId, StudentId: studentId}
	return DB.Self.Create(&bind).Error
}

// 删除绑定关系，返回受影响的行数
func DeleteUserStudent(userId uint64) (int64, error) {
	result := DB.Self.Where("user_id = ?", userId).Delete(&UserStudent{})
	return result.RowsAffected, result.Error
}
