package middleware

import (
	"net/http"
	"strings"

	"github.com/Lenoud/gin-demo/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 验证 JWT 的中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Authorization 头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求头缺少 Authorization"})
			c.Abort() // 阻止继续执行
			return
		}

		// 2. 必须是 "Bearer <token>" 格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization 格式错误"})
			c.Abort()
			return
		}

		// 3. 解析 token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 Token: " + err.Error()})
			c.Abort()
			return
		}

		// 4. 将用户信息存入上下文，方便后续 handler 使用
		c.Set("user_id", claims["user_id"])
		c.Set("is_admin", claims["is_admin"])

		// 放行
		c.Next()
	}
}
