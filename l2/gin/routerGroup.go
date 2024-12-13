package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 简单的路由组: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// 简单的路由组: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

func loginEndpoint(c *gin.Context) {

	c.String(http.StatusOK, c.Request.URL.String()+":hello/login")
}
func submitEndpoint(c *gin.Context) {
	c.String(http.StatusOK, c.Request.URL.String()+":hello/submit")
}
func readEndpoint(c *gin.Context) {
	c.String(http.StatusOK, c.Request.URL.String()+":hello/read")
}
