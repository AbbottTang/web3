//func main() {
//	r := gin.Default()
//
//	// Ping handler
//	r.GET("/ping", func(c *gin.Context) {
//		c.String(200, "pong")
//	})
//
//	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
//}

package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	m := autocert.Manager{
		//自动接受证书颁发机构的服务条款。
		Prompt: autocert.AcceptTOS,
		//指定证书缓存的目录，以便在证书更新时能够重用之前的证书。
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		///指定证书缓存的目录，以便在证书更新时能够重用之前的证书。
		Cache: autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}

//您的代码示例展示了一个使用Gin框架和Gin的autotls库来自动配置TLS并启动一个HTTPS服务器的简单应用。这里是对您代码的详细解释以及一些潜在的注意事项：
//
//代码解释
//‌包引入‌
//
//引入了log包用于记录日志。
//引入了github.com/gin-gonic/autotls包，这是一个Gin框架的扩展，用于自动处理TLS证书的申请和更新。
//引入了github.com/gin-gonic/gin包，这是Gin框架的核心包，用于构建HTTP服务器和处理请求。
//‌主函数main‌
//
//创建一个默认的Gin路由器r，它包含了一些默认的中间件，如Logger和Recovery。
//为路由器添加一个处理函数，当访问/ping路径时，会返回一个状态码为200的响应，以及文本内容"pong"。
//调用autotls.Run函数来启动服务器，并自动配置TLS。这个函数接受两个参数：Gin路由器和要监听的主机名（在这个例子中是"example1.com"和"example2.com"）。autotls.Run会处理证书的获取和更新，以及HTTPS服务器的启动。
//如果服务器启动失败，autotls.Run会返回一个错误，这个错误会被log.Fatal记录并导致程序退出。
//注意事项
//‌域名和证书‌
//
//确保"example1.com"和"example2.com"是您拥有或有权使用的域名。
//autotls通常使用Let's Encrypt等证书颁发机构来自动获取证书。确保您的域名已经正确解析，并且服务器可以从外部访问到，以便证书颁发机构可以验证域名的所有权。
//‌错误处理‌
//
//在生产环境中，您可能希望更细致地处理错误，而不是简单地使用log.Fatal。例如，您可以记录错误详情，并尝试重启服务器或执行其他恢复操作。
//‌安全性‌
//
//确保您的服务器和代码都是安全的。除了使用TLS来保护数据传输外，还应该关注代码中的其他安全漏洞，如SQL注入、跨站脚本攻击（XSS）等。
//‌配置和定制‌
//
//autotls提供了一些配置选项来定制TLS的行为，例如指定证书存储位置、设置证书更新的频率等。您可以查阅autotls的文档来了解更多配置选项。
//‌依赖管理‌
//
//确保您的项目中包含了所有必要的依赖项，并且它们的版本是兼容的。使用Go的模块系统（如go.mod）来管理依赖项是一个好习惯。
//‌日志记录‌
//
//考虑在生产环境中使用更专业的日志记录库，如logrus或zap，它们提供了更丰富的日志记录功能和更好的性能。
//‌测试‌
//
//在将代码部署到生产环境之前，确保对其进行充分的测试，包括单元测试、集成测试和性能测试。
//综上所述，您的代码是一个很好的起点，用于构建一个使用Gin框架和autotls的HTTPS服务器。根据您的具体需求，您可能需要对代码进行一些调整和优化。
