package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始的时间
		t := time.Now()

		// 设置 example 变量
		// 在上下文中设置一个键值对，键为 "example"，值为 "12345"
		// 这个值可以在后续的中间件或路由处理函数中通过 c.Get("example") 获取
		c.Set("example", "12345")

		// 请求前
		// 调用 c.Next() 将会使请求继续传递到下一个中间件或最终的路由处理函数
		c.Next()

		// 请求后
		// 请求处理完成后，计算请求处理的耗时
		latency := time.Since(t)
		log.Print(latency)

		// 获取并打印HTTP响应状态码
		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	// 使用 Logger 中间件
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		// 从上下文中获取 "example" 值，并断言为 string 类型
		example := c.MustGet("example").(string)

		// 打印："12345"
		log.Println(example)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
