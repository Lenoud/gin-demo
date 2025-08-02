package student

import (
	"errors"
	"strconv"

	"github.com/Lenoud/gin-demo/controller"
	"github.com/Lenoud/gin-demo/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AddStudent(c *gin.Context) {
	var student model.StudentInfo
	if err := c.ShouldBindJSON(&student); err != nil {
		controller.SendResponse(c, 400, "接收数据失败！", nil)
		return
	}
	if err := student.Create(); err != nil {
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
	student := &model.StudentInfo{Id: id}
	if err := student.Delete(); err != nil {
		controller.SendResponse(c, 500, "删除失败", err.Error())
		return
	}
	controller.SendResponse(c, 200, "删除成功", gin.H{"deleted_id": id})
}

func UpdateStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		controller.SendResponse(c, controller.CodeParamError, "无效的ID格式，请输入数字", nil)
		return
	}

	// 从请求体获取要更新的信息
	var updateData model.StudentInfo
	if err := c.ShouldBindJSON(&updateData); err != nil {
		controller.SendResponse(c, controller.CodeParamError, "接收更新数据失败："+err.Error(), nil)
		return
	}

	// 初始化学生对象并设置ID
	student := model.StudentInfo{Id: id}
	// 调用模型层的Update方法
	if err := student.Update(&updateData); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			controller.SendResponse(c, controller.CodeNotFound, "未找到该学生", nil)
		} else {
			controller.SendResponse(c, controller.CodeServerError, "更新失败："+err.Error(), nil)
		}
		return
	}

	// 返回更新结果
	controller.SendResponse(c, controller.CodeSuccess, "更新成功",
		gin.H{
			"updated_id": id,
			"student":    student,
		})
}

func GetStudents(c *gin.Context) {
	var student model.StudentInfo
	students, err := student.GetAll()
	if err != nil {
		controller.SendResponse(c, 500, "查询失败: "+err.Error(), nil)
		return
	}

	controller.SendResponse(c, 200, "查询成功!",
		gin.H{
			"count": len(students),
			"list":  students,
		})
}

// func UpdateStudent(c *gin.Context) {
// 	idStr := c.Param("id")
// 	//转换ID格式
// 	id, err := strconv.ParseUint(idStr, 10, 64)
// 	if err != nil {
// 		controller.SendResponse(c, controller.CodeParamError, "无效的ID格式，请输入数字", nil)
// 		return
// 	}

// 	//从请求体获取要更新的信息
// 	var updateData model.StudentInfo
// 	if err := c.ShouldBindJSON(&updateData); err != nil {
// 		controller.SendResponse(c, controller.CodeParamError, "接收更新数据失败："+err.Error(), nil)
// 		return
// 	}

// 	// 执行更新操作
// 	// 先查询学生是否存在
// 	var student model.StudentInfo
// 	if err := model.DB.Self.First(&student, id).Error; err != nil {
// 		controller.SendResponse(c, controller.CodeNotFound, "未找到该学生", nil)
// 		return
// 	}

// 	// 执行更新（只更新非零值字段，或指定字段）
// 	result := model.DB.Self.Model(&student).Updates(updateData)
// 	if result.Error != nil {
// 		controller.SendResponse(c, controller.CodeServerError, "更新失败："+result.Error.Error(), nil)
// 		return
// 	}

// 	// 6. 返回更新结果
// 	controller.SendResponse(c, controller.CodeSuccess, "更新成功", gin.H{
// 		"updated_id": id,
// 		"student":    student,
// 	})
// }

// func GetStudents(c *gin.Context) {
// 	var students []model.StudentInfo
// 	result := model.DB.Self.Find(&students)
// 	if result.Error != nil {
// 		controller.SendResponse(c, 500, "查询失败: "+result.Error.Error(), nil)
// 		return
// 	}

// 	controller.SendResponse(c, 0, "查询成功!", gin.H{
// 		"count": len(students), // 增加记录总数
// 		"list":  students,      // 返回学生列表
// 	})
// }

// func DelStudent(c *gin.Context) {
// 	idStr := c.Param("id")
// 	var student model.StudentInfo
// 	//将字符串ID转换为uint64类型
// 	id, err := strconv.ParseUint(idStr, 10, 64)
// 	if err != nil {
// 		controller.SendResponse(c, 400, "无效的ID格式，请输入数字", nil)
// 		return
// 	}

// 	result := model.DB.Self.Delete(&student, id)
// 	if result.Error != nil {
// 		controller.SendResponse(c, 500, "删除失败", result.Error.Error())
// 		return
// 	}
// 	if result.RowsAffected == 0 {
// 		controller.SendResponse(c, 404, "未找到ID", nil)
// 		return
// 	}
// 	controller.SendResponse(c, 200, "删除成功", gin.H{"deleted_id": id})
// }
