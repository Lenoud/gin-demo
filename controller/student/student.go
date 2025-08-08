// controller/student/student.go
package student

import (
	"strconv"

	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/model"
	"github.com/Lenoud/gin-demo/service"
	"github.com/gin-gonic/gin"
)

func AddStudent(c *gin.Context) {
	var student model.StudentInfo
	if err := c.ShouldBindJSON(&student); err != nil {
		controller.SendResponse(c, 400, "接收数据失败！", nil)
		return
	}

	if err := service.AddStudent(&student); err != nil {
		controller.SendResponse(c, 500, "写入数据失败！", nil)
		return
	}

	controller.SendResponse(c, 200, "创建学生信息成功！", gin.H{"user": student})
}

func DelStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, 400, "无效的ID格式，请输入数字", nil)
		return
	}

	if err := service.DelStudent(id); err != nil {
		controller.SendResponse(c, 400, "删除失败", err.Error())
		return
	}

	controller.SendResponse(c, 200, "删除成功", gin.H{"deleted_stu_id": id})
}

// 处理HTTP请求的UpdateStudent控制器方法
func UpdateStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, controller.CodeParamError, "无效的ID格式，请输入数字", nil)
		return
	}

	var updateData model.StudentInfo
	if err := c.ShouldBindJSON(&updateData); err != nil {
		controller.SendResponse(c, controller.CodeParamError, "接收更新数据失败："+err.Error(), nil)
		return
	}

	updatedStudent, err := service.UpdateStudent(id, &updateData)
	if err != nil {
		controller.SendResponse(c, controller.CodeServerError, "更新失败："+err.Error(), nil)
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
	students, err := service.GetAllStudents()
	if err != nil {
		controller.SendResponse(c, controller.CodeServerError, "查询失败: "+err.Error(), nil)
		return
	}

	controller.SendResponse(c, controller.CodeSuccess, "查询成功!",
		gin.H{
			"count": len(students),
			"list":  students,
		})
}
