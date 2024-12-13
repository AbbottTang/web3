package main

import (
	"fmt"
	"time"
)

//http://localhost:8080/ping
//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
//}

// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON。
// http://localhost:8080/someJSON
//func main() {
//	r := gin.Default()
//
//	r.GET("/someJSON", func(c *gin.Context) {
//		data := map[string]interface{}{
//			"lang": "GO语言",
//			"tag":  "<br>",
//		}
//
//		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
//		c.AsciiJSON(http.StatusOK, data)
//	})
//
//	// 监听并在 0.0.0.0:8080 上启动服务
//	r.Run(":8080")
//}

// 使用 LoadHTMLGlob() 或者 LoadHTMLFiles()
// http://localhost:8080/index
//func main() {
//	router := gin.Default()
//	router.LoadHTMLGlob("templates/*")
//	router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
//	router.GET("/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "template1.html", gin.H{
//			"title": "Main website",
//		})
//	})
//	router.Run(":8080")
//}

// 使用不同目录下名称相同的模板
// http://localhost:8080/posts/index
// http://localhost:8080/users/index
//func main() {
//	router := gin.Default()
//	router.LoadHTMLGlob("templates/**/*")
//	router.GET("/posts/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
//			"title": "Posts",
//		})
//	})
//	router.GET("/users/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
//			"title": "Users",
//		})
//	})
//	router.Run(":8080")
//}
//自定义模板渲染器
//你可以使用自定义的 html 模板渲染

//	func main() {
//		router := gin.Default()
//		html := template.Must(template.ParseFiles("file1", "file2"))
//		router.SetHTMLTemplate(html)
//		router.Run(":8080")
//	}
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

// 自定义分隔符
// 你可以使用自定义分隔
//func main() {
//	router := gin.Default()
//	router.Delims("{[{", "}]}")
//	router.SetFuncMap(template.FuncMap{
//		"formatAsDate": formatAsDate,
//	})
//	router.LoadHTMLFiles("templates/testdata/raw.tmpl")
//
//	router.GET("/raw", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "raw.tmpl", gin.H{
//			"now": time.Date(2017, 0o7, 0o1, 0, 0, 0, 0, time.UTC),
//		})
//	})
//
//	router.Run(":8080")
//}
