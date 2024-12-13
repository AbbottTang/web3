package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// c.Query和c.DefaultQuery是用于从HTTP请求的查询参数（即URL中?后面的部分）中提取值的方法。
// 这些方法都作用于Gin的上下文对象c，该对象代表当前的HTTP请求和响应。
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}
