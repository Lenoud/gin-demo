// controller/student/student.go
package student

import (
	"strconv"

	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/model"
	"github.com/Lenoud/gin-demo/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddStudent(c *gin.Context) {
	zap.L().Info("收到添加学生请求")
	var student model.StudentInfo
	if err := c.ShouldBindJSON(&student); err != nil {
		controller.SendResponse(c, 400, "接收数据失败！", nil)
		zap.L().Warn("接收学生数据失败", zap.String("error", err.Error()))
		return
	}

	if err := service.AddStudent(&student); err != nil {
		controller.SendResponse(c, 500, "写入数据失败！", nil)
		zap.L().Error("写入学生数据失败", zap.String("error", err.Error()))
		return
	}

	controller.SendResponse(c, 200, "创建学生信息成功！", gin.H{"user": student})
}

func DelStudent(c *gin.Context) {
	zap.L().Info("收到删除学生请求", zap.String("studentId", c.Param("id")))
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, 400, "无效的ID格式，请输入数字", nil)
		zap.L().Warn("无效的学生ID格式", zap.String("id", idStr), zap.String("error", err.Error()))
		return
	}

	if err := service.DelStudent(id); err != nil {
		controller.SendResponse(c, 400, "删除失败", err.Error())
		zap.L().Error("删除学生失败", zap.Uint64("studentId", id), zap.String("error", err.Error()))
		return
	}

	controller.SendResponse(c, 200, "删除成功", gin.H{"deleted_stu_id": id})
}

// 处理HTTP请求的UpdateStudent控制器方法
func UpdateStudent(c *gin.Context) {
	zap.L().Info("收到更新学生请求", zap.String("studentId", c.Param("id")))
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, controller.CodeParamError, "无效的ID格式，请输入数字", nil)
		zap.L().Warn("无效的学生ID格式", zap.String("id", idStr), zap.String("error", err.Error()))
		return
	}

	var updateData model.StudentInfo
	if err := c.ShouldBindJSON(&updateData); err != nil {
		controller.SendResponse(c, controller.CodeParamError, "接收更新数据失败："+err.Error(), nil)
		zap.L().Warn("接收更新学生数据失败", zap.String("error", err.Error()))
		return
	}

	updatedStudent, err := service.UpdateStudent(id, &updateData)
	if err != nil {
		controller.SendResponse(c, controller.CodeServerError, "更新失败："+err.Error(), nil)
		zap.L().Error("更新学生失败", zap.Uint64("studentId", id), zap.String("error", err.Error()))
		return
	}

	controller.SendResponse(c, controller.CodeSuccess, "更新成功",
		gin.H{
			"updated_id": id,
			"student":    updatedStudent,
		})
}

// 处理HTTP请求的GetStudents控制器方法
func GetStudents(c *gin.Context) {
	zap.L().Info("收到获取学生列表请求")
	students, err := service.GetAllStudents()
	if err != nil {
		controller.SendResponse(c, controller.CodeServerError, "查询失败: "+err.Error(), nil)
		zap.L().Error("查询学生列表失败", zap.String("error", err.Error()))
		return
	}

	controller.SendResponse(c, controller.CodeSuccess, "查询成功!",
		gin.H{
			"count": len(students),
			"list":  students,
		})
}
