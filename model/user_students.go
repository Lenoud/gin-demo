package model

type UserStudent struct {
	UserId    uint64 `gorm:"primaryKey;autoIncrement:false" json:"user_id"`
	StudentId uint64 `gorm:"primaryKey;autoIncrement:false" json:"student_id"`
}

func (UserStudent) TableName() string {
	return "user_students"
}
