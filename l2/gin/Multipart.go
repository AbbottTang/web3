package main

import "github.com/gin-gonic/gin"

// 这里，User 是一个类型为 string 的字段，它配备了两个标签：form 和 binding。
//
// ‌form:"user"‌
//
// form 标签指明了该字段应从HTTP请求的表单数据中获取数据，并且使用的表单字段名为 "user"。
// 当Gin框架处理HTTP请求并尝试将数据绑定到结构体时，它会查找名为 "user" 的表单字段，并将其值赋给 User 字段。
// ‌binding:"required"‌
//
// binding 标签用于定义字段的绑定和验证规则。
// 在此例中，"required" 表明该字段是必填的。如果HTTP请求中未包含 "user" 字段，或该字段的值为空，Gin将返回一个错误，指出该字段是必需的。
// 这种标签机制使得Gin框架能够轻松地处理HTTP请求中的数据，并将其映射到Go语言的结构体中，同时执行必要的验证。这对于构建RESTful API和Web应用程序非常有用，因为它简化了数据解析、验证和错误处理的过程。
//type LoginForm struct {
//	User     string `form:"user" binding:"required"`
//	Password string `form:"password" binding:"required"`
//}

//	func main() {
//		router := gin.Default()
//		router.POST("/login", func(c *gin.Context) {
//			// 你可以使用显式绑定声明绑定 multipart form：
//			// c.ShouldBindWith(&form, binding.Form)
//			// 或者简单地使用 ShouldBind 方法自动绑定：
//			var form LoginForm
//			// 在这种情况下，将自动选择合适的绑定
//			if c.ShouldBind(&form) == nil {
//				if form.User == "user" && form.Password == "password" {
//					c.JSON(200, gin.H{"status": "you are logged in"})
//				} else {
//					c.JSON(401, gin.H{"status": "unauthorized"})
//				}
//			}
//		})
//		router.Run(":8080")
//	}
func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}
