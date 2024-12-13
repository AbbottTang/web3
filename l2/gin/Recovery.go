package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

// 自定义中间件示例
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在这里实现认证逻辑
		// 如果认证失败，可以使用 c.Abort() 终止请求
		// 例如: if !authenticated { c.AbortWithStatus(401) }
		// 这里只是示例，所以直接通过
		gin.BasicAuth(gin.Accounts{
			"foo":    "bar",
			"austin": "1234",
			"lena":   "hello2",
			"manu":   "4321",
		})
	}
}

// 示例处理函数
func benchEndpoint(c *gin.Context) {
	// 处理基准测试请求
	c.String(http.StatusOK, c.Request.URL.String()+":hello/benchEndpoint")
}

func analyticsEndpoint(c *gin.Context) {
	// 处理分析请求
	c.String(http.StatusOK, c.Request.URL.String()+":hello/analyticsEndpoint")
}
func loginEndpoint1(c *gin.Context) {
	// 处理登录请求
	c.String(http.StatusOK, c.Request.URL.String()+":hello/loginEndpoint1")

}

func submitEndpoint1(c *gin.Context) {
	// 处理提交请求
	c.String(http.StatusOK, c.Request.URL.String()+":hello/submitEndpoint1")

}

func readEndpoint1(c *gin.Context) {
	// 处理读取请求
	c.String(http.StatusOK, c.Request.URL.String()+":hello/readEndpoint1")

}

// 另一个自定义中间件示例，用于记录基准测试
func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在这里记录基准测试相关的日志
		// ...
		// 继续处理请求
		gin.DefaultWriter = io.MultiWriter(os.Stdout)
	}
}

func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局中间件
	r.Use(gin.Logger())   // Logger 中间件将日志写入 gin.DefaultWriter
	r.Use(gin.Recovery()) // Recovery 中间件会 recover 任何 panic

	// 添加一个处理基准测试的路由
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 认证路由组
	authorized := r.Group("/authorized")
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的 AuthRequired() 中间件
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint1)
		authorized.POST("/submit", submitEndpoint1)
		authorized.POST("/read", readEndpoint1)

		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
