package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/gopher.png", "./resources/gopher (1).png")

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}
