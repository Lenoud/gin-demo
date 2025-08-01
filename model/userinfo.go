package model


// UserJson 结构体，存储用户信息
type UserJson struct {
	// binding:"required" 代表必填项 没填写报错
    Id       int    `json:"id"`
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Age      int    `json:"age"`  // 这里定义为 int 类型
}