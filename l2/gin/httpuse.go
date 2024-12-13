package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run()
	// router.Run(":3000") hardcode 端口号
}
func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"getting": "someGet"})
}
func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"getting": "somePost"})
}
func putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"getting": "somePut"})
}
func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"getting": "someDelete"})
}
func patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"getting": "somePatch"})
}
func head(c *gin.Context) {
	// 设置一些自定义的响应头
	c.Header("X-Custom-Header", "headValue")
	c.JSON(http.StatusOK, gin.H{"getting": "someHead"})
}
func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"getting": "someOptions"})
}
