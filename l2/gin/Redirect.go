package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})
	route.POST("/test", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})
	route.Any("/foo", fooPage)
	route.GET("/test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		route.HandleContext(c)
	})
	route.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	route.Run(":8080")
}
func fooPage(c *gin.Context) {

	c.String(200, "/foo")
}
