package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORS 跨域
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许携带证书信息（如cookies）
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// 允许的HTTP方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// 允许的HTTP头部
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		c.Next()
	}
}
