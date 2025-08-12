package controller

const (
	CodeSuccess     = 200 // 成功
	CodeParamError  = 400 // 参数错误
	CodeNotFound    = 404 // 资源不存在
	CodeServerError = 500 // 服务器错误
)

// 响应状态码定义（RespCode = Response Code）
// 注：自定义状态码建议避开 HTTP 原生状态码语义冲突（如 401/403 等可复用标准含义）
// const (
// 	RespCodeSuccess     = 200 // 操作成功（通用成功状态）
// 	RespCodeParamError  = 400 // 参数错误（请求参数格式/校验失败）
// 	RespCodeUnauthorized = 401 // 未授权（未登录或 token 失效）
// 	RespCodeForbidden   = 403 // 权限不足（无操作权限）
// 	RespCodeNotFound    = 404 // 资源不存在（请求的资源未找到）
// 	RespCodeServerError = 500 // 服务器错误（后端逻辑/数据库异常）
// 	RespCodeBusy        = 503 // 服务繁忙（如限流、降级场景）
// 	RespCodeConflict    = 409 // 资源冲突（如重复创建唯一资源）
// )
