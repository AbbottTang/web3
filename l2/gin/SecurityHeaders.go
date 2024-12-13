// 安全页眉
// 使用安全标头保护网络应用程序免受常见安全漏洞的攻击非常重要。
// 本示例将向您展示如何在 Gin 应用程序中添加安全标头，以及如何避免与主机标头注入相关的攻击（SSRF、开放重定向）。
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//r := gin.Default()：创建一个带有默认中间件（如Logger和Recovery）的Gin引擎。
	r := gin.Default()

	expectedHost := "localhost:8080"
	//使用r.Use添加一个全局中间件，该中间件会在每个请求之前执行。
	//在中间件中，首先检查请求的Host头是否与预期的expectedHost相匹配。
	//如果不匹配，则使用c.AbortWithStatusJSON方法中断请求，并返回400 Bad Request响应和错误消息。
	//然后，设置一系列安全相关的HTTP头，以增强应用的安全性。
	//这些头包括防止点击劫持、内容安全策略、XSS保护、HSTS、引用策略、内容类型选项和权限策略等。
	// Setup Security Headers
	r.Use(func(c *gin.Context) {
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
