# Gin Demo 项目

一个基于 Gin 框架的 Go 后端示例项目，集成用户注册、登录、学生与成绩管理等功能，支持 JWT 认证与跨域处理，适用于学习和快速开发 RESTful API 服务。

## 技术栈

- Go 1.24.5+
- Gin Web 框架
- MySQL（用户与业务数据存储）
- Zap（高性能日志库）
- JWT（用户认证）
- Redis（可选，缓存支持）

## 目录结构

```text
gin-demo/
├── cmd/
│   └── api/
│       └── main.go         # API 服务入口与启动初始化
├── configs/
│   └── config.yaml         # 运行配置文件
├── internal/               # 本项目内部使用的 Go 包
│   ├── config/             # 配置加载与环境变量覆盖
│   ├── controller/         # HTTP 请求处理与统一响应
│   ├── middleware/         # CORS、JWT 等中间件
│   ├── model/              # 数据模型与数据库操作
│   ├── router/             # 路由注册
│   ├── service/            # 业务逻辑
│   └── utils/              # JWT 工具与日志初始化
├── examples/
│   ├── basic/              # 独立的 Gin 入门示例
│   └── webform/            # 登录与注册表单示例
├── bin/                    # 本地构建产物，已忽略
├── logs/                   # 运行时日志，已忽略
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

`bin/` 和 `logs/` 在构建或运行时生成。`internal/` 使用 Go 的内部包机制，包只能由其父目录范围内的代码导入。

### 目录约定

- 新增可执行程序放到 `cmd/<程序名>/`，入口负责配置、依赖初始化和启动。
- HTTP 参数校验与响应放到 `internal/controller/`，业务处理放到 `internal/service/`，数据库读写放到 `internal/model/`。
- YAML 等运行配置放到 `configs/`，配置加载代码放到 `internal/config/`。
- 示例程序放到 `examples/<示例名>/`，独立运行。
- Go 测试与被测代码放在同一目录，文件名使用 `_test.go` 后缀。

## 安装与部署

以下命令均在包含 `go.mod` 的项目根目录执行。配置文件和日志的相对路径以运行时工作目录为基准。

1. **环境准备**
   - 安装 Go 1.24.5 及以上版本
   - 配置 MySQL 数据库，并在 `configs/config.yaml` 中填写连接信息

2. **依赖安装**
   ```bash
   go mod tidy
   ```

3. **配置修改**
   - 编辑 `configs/config.yaml`，设置数据库、端口等参数

4. **启动服务**
   ```bash
   go run ./cmd/api
   ```
   或
   ```bash
   mkdir -p bin
   go build -o bin/gin-demo ./cmd/api
   ./bin/gin-demo
   ```

5. **开发验证**
   ```bash
   go test ./...
   go vet ./...
   ```

### 运行入门示例

原根目录 `server.go` 中的 Gin 示例已整理到 `examples/basic/`，可独立运行，无需连接数据库：

```bash
go run ./examples/basic
```

示例监听 `:38080`，访问 `http://127.0.0.1:38080/ping` 可获得 `online` 响应。

### 运行网页表单示例

在项目根目录运行 `go run ./examples/webform`，然后访问 `http://127.0.0.1:38081`。
示例演示注册表单和登录校验。账号只在进程内保存，重启后需要重新注册。

## 核心模块说明

- **configs/ 与 internal/config/**：分别存放 YAML 配置文件与配置加载代码，支持环境变量覆盖。

- **internal/controller/**：路由控制器，按功能分模块（如 login、register、student、score），负责参数校验与响应封装。

- **internal/middleware/**：公共中间件，包括 CORS 跨域处理和 JWT 认证。

- **internal/model/**：数据模型定义及数据库操作方法。

- **internal/service/**：业务逻辑实现，处理具体业务流程（如用户注册、登录、学生成绩管理）。

- **internal/utils/**：工具类，包含 JWT 生成校验、日志初始化等。

## API 示例

### 用户注册

```http
POST /api/register
Content-Type: application/json
Authorization: Bearer <token>

{
  "username": "testuser",
  "password": "123456",
  "email": "test@example.com"
}
```

### 用户登录

```http
POST /api/auth/login
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

- 具体接口见 `internal/controller/student/student.go` 和 `internal/controller/score/score.go`，支持学生信息与成绩的增删查改。

## 注意事项

- 日志文件默认输出至 `logs/app.log`，可在配置文件中调整。
- JWT 密钥请在 `configs/config.yaml` 中设置；Token 有效期由 `internal/utils/jwt.go` 中的生成逻辑控制。
- 跨域策略可在 `internal/middleware/cors.go` 中自定义。
- 数据库表结构需与 `internal/model/` 下的数据模型一致。
- 如需使用 Redis，请补充相关配置与依赖。

## 相关链接

- [Gin 官方文档](https://gin-gonic.com/docs/)
- [Zap 日志库](https://github.com/uber-go/zap)
- [JWT Go 实现](https://github.com/dgrijalva/jwt-go)

---
如有问题或建议，欢迎提交 Issue 或 PR。
