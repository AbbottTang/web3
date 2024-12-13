/*
*
net/http服务器 中的一个基本概念是处理程序。处理程序是实现接口的对象 http.Handler。
编写处理程序的常用方法是使用http.HandlerFunc具有适当签名的函数上的适配器。
充当处理程序的函数以 a http.ResponseWriter和 ahttp.Request作为参数。
响应编写器用于填写 HTTP 响应。这里我们的简单响应只是“hello\n”。
该处理程序通过读取所有 HTTP 请求标头并将它们回显到响应主体中，执行一些更复杂的事情。
我们使用便捷函数在服务器路由上注册处理程序 。它在包中http.HandleFunc设置默认路由器net/http并将函数作为参数。
ListenAndServe最后，我们使用端口和处理程序调用。nil告诉它使用我们刚刚设置的默认路由器。

在后台运行服务器。go run http-servers.go &
访问/hello路线。curl localhost:8090/hello
*/
package main

import (
	"fmt"
	"net/http"
)

// hello 函数是一个HTTP处理器，当访问/hello路径时会被调用。
// 它向HTTP响应中写入"hello\n"字符串。
func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

// headers 函数是另一个HTTP处理器，当访问/headers路径时会被调用。
// 它遍历HTTP请求中的所有头部信息，并将它们以"name: value"的格式写入HTTP响应中。
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// main 函数是程序的入口点。
// 它注册了两个HTTP处理器：一个处理/hello路径，另一个处理/headers路径。
// 然后，它启动HTTP服务器，监听本地的8090端口。
func main() {
	// 注册/hello路径的处理器为hello函数
	http.HandleFunc("/hello", hello)
	// 注册/headers路径的处理器为headers函数
	http.HandleFunc("/headers", headers)
	// 启动HTTP服务器，监听本地的8090端口，并传入nil作为默认的处理器（因为我们已经注册了具体的路径处理器）
	http.ListenAndServe(":8090", nil)
}
