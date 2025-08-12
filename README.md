# Gin Demo 项目

一个基于 Gin 框架的 Go 后端示例项目，集成用户注册、登录、学生与成绩管理等功能，支持 JWT 认证与跨域处理，适用于学习和快速开发 RESTful API 服务。

## 技术栈

- Go 1.21+
- Gin Web 框架
- MySQL（用户与业务数据存储）
- Zap（高性能日志库）
- JWT（用户认证）
- Redis（可选，缓存支持）

## 目录结构

```
├── config/         # 配置文件与加载逻辑
├── controller/     # 路由控制器，处理请求与响应
├── logs/           # 项目日志输出
├── main.go         # 项目入口
├── middleware/     # 中间件（如CORS、JWT认证）
├── model/          # 数据模型与数据库操作
├── router/         # 路由注册
├── server.go       # 服务启动逻辑
├── service/        # 业务逻辑层
├── utils/          # 工具类（如JWT、日志）
```

## 安装与部署

1. **环境准备**
   - 安装 Go 1.21 及以上版本
   - 配置 MySQL 数据库，并在 `config/config.yaml` 中填写连接信息

2. **依赖安装**
   ```bash
   go mod tidy
   ```

3. **配置修改**
   - 编辑 `config/config.yaml`，设置数据库、端口等参数

4. **启动服务**
   ```bash
   go run main.go
   ```
   或
   ```bash
   go build -o gin-demo
   ./gin-demo
   ```

## 核心模块说明

- **config/**  
  配置加载与管理，支持 YAML 格式。

- **controller/**  
  路由控制器，按功能分模块（如 login、register、student、score），负责参数校验与响应封装。

- **middleware/**  
  公共中间件，包括 CORS 跨域处理和 JWT 认证。

- **model/**  
  数据模型定义及数据库操作方法。

- **service/**  
  业务逻辑实现，处理具体业务流程（如用户注册、登录、学生成绩管理）。

- **utils/**  
  工具类，包含 JWT 生成校验、日志初始化等。

## API 示例

### 用户注册

```http
POST /register
Content-Type: application/json

{
  "username": "testuser",
  "password": "123456",
  "email": "test@example.com"
}
```

### 用户登录

```http
POST /login
Content-Type: application/json

{
  "username": "testuser",
  "password": "123456"
}
```

返回示例：

```json
{
  "token": "jwt_token_string",
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "is_admin": false
  }
}
```

### 学生与成绩管理

- 具体接口见 `controller/student/student.go` 和 `controller/score/score.go`，支持学生信息与成绩的增删查改。

## 注意事项

- 日志文件默认输出至 `logs/app.log`，可在配置文件中调整。
- JWT 密钥与过期时间请在 `config/config.yaml` 中设置。
- 跨域策略可在 `middleware/cors.go` 中自定义。
- 数据库表结构需与 `model/` 下的数据模型一致。
- 如需使用 Redis，请补充相关配置与依赖。

## 相关链接

- [Gin 官方文档](https://gin-gonic.com/docs/)
- [Zap 日志库](https://github.com/uber-go/zap)
- [JWT Go 实现](https://github.com/dgrijalva/jwt-go)

---
如有问题或建议，欢迎提交 Issue 或 PR。