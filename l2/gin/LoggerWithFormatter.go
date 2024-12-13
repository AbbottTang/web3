package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.New()
	// LoggerWithFormatter 中间件会写入日志到 gin.DefaultWriter
	// 默认 gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.Run(":8080")
}

//使用router.Use()方法添加中间件。在这里，您添加了两个中间件：
//‌自定义日志中间件‌：通过gin.LoggerWithFormatter函数，您传入了一个自定义的格式化函数。这个函数会根据传入的gin.LogFormatterParams参数生成一个格式化的字符串，该字符串包含了客户端IP、时间戳、HTTP方法、请求路径、协议版本、状态码、延迟、用户代理和错误信息（如果有的话）。这个格式化字符串随后被写入到gin.DefaultWriter，默认情况下它是os.Stdout。
//‌恢复中间件‌：gin.Recovery()是一个内置的Gin中间件，用于恢复从任何panic中，如果有的话，它将写入一个500内部服务器错误响应。
//为路由器添加一个处理函数，当访问/ping路径时，会返回一个状态码为200的响应和文本内容"pong"。
