package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 如果你想要以指定的格式（例如 JSON，key values 或其他格式）记录信息，则可以使用 gin.DebugPrintRouteFunc 指定格式。
// 在下面的示例中，我们使用标准日志包记录所有路由，但你可以使用其他满足你需求的日志工具。
func main() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
