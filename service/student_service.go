// service/student_service.go
package service

import (
	"github.com/Lenoud/gin-demo/model"
	"go.uber.org/zap"
)

// 添加学生
func AddStudent(student *model.StudentInfo) error {
	zap.L().Info("尝试添加学生", zap.String("studentName", student.StuName))
	// 这里可以添加业务校验逻辑
	// 例如：if student.StrAge < 6 { return errors.New("年龄不能小于6岁") }

	return student.Create()
}

// 删除学生
func DelStudent(id uint64) error {
	zap.L().Info("尝试删除学生", zap.Uint64("studentId", id))
	student := &model.StudentInfo{Id: id}
	return student.Delete()
}

// 更新学生
func UpdateStudent(id uint64, updateData *model.StudentInfo) (*model.StudentInfo, error) {
	zap.L().Info("尝试更新学生信息", zap.Uint64("studentId", id))
	student := model.StudentInfo{Id: id}
	if err := student.Update(updateData); err != nil {
		return nil, err
	}
	return &student, nil
}

// 获取所有学生
func GetAllStudents() ([]model.StudentInfo, error) {
	zap.L().Info("尝试获取所有学生信息")
	var student model.StudentInfo
	return student.GetAll()
}
