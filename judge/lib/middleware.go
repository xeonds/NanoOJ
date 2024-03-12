package lib

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Gin中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		log.Print(latency)
	}
}

func JWTMiddleware(setUserToken func(c *gin.Context, userClaim UserClaim)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.AbortWithStatus(401)
			return
		}
		token := c.GetHeader("Authorization")
		parsed, err := ParseToken(token)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		if parsed.Valid() != nil {
			c.AbortWithStatus(401)
			return
		}
		if setUserToken != nil {
			// 设置用户信息字段
			setUserToken(c, *parsed)
		}
		c.Next()
	}
}

// 验证用户权限等级
// 权限位于[permLo, permHi]之间则为合理
func AuthMiddleware(permLo, permHi int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.AbortWithStatus(401)
		}
		token := c.GetHeader("Authorization")
		parsed, err := ParseToken(token)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		if parsed.Valid() != nil {
			c.AbortWithStatus(401)
			return
		}
		// TODO: query from db in case of priviledge delay
		if parsed.Permission < permLo || parsed.Permission > permHi {
			c.AbortWithStatus(403)
			return
		}
		c.Next()
	}
}

// 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// finish this golang cors middleware
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Next()
	}
}
