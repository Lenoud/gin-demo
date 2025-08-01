package model

// User 结构体，存储用户信息
type User struct {
	Id       int
	Username string
	Password string
	// 以后可以加更多字段，比如 Email、CreatedAt 等
}

var Users = map[string]User{}

type UserJson struct {
	// binding:"required" 代表必填项 没填写报错
	Id       int    `json:"id" binding:"required"`  
	Username string `json:"username"`
	Password string `json:"password"`
	Age      string `json:"age"`
}
