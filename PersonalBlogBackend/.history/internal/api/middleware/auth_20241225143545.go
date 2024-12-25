package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// 先尝试从URL参数获取token
			token := c.Query("Authorization")
			if token == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})
				c.Abort()
				return
			}
			authHeader = "Bearer " + token
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "认证格式无效"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的认证令牌"})
			c.Abort()
			return
		}

		// 将用户ID转换为uint
		userID := uint(claims["user_id"].(float64))
		c.Set("userID", userID)
		c.Next()
	}
}
