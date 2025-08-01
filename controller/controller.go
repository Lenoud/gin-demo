package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// 统一标准化返回 json
/*
{
"code":"",  应用的状态码
"data":"",  返回的数据
"message":""  提示信息 对状态码说明
}
*/

type Response struct{
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SendResponse(c *gin.Context,code int ,message string, data interface{}){
	rsp:= Response{
		Code: code,
		Message: message,
		Data: data,
	}
	c.JSON(http.StatusOK,rsp)
}