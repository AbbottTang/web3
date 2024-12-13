/*
*
在前面的例子中，我们研究了如何设置一个简单的 HTTP 服务器context.Context。
HTTP 服务器对于演示控制取消的用法很有用。
AContext跨 API 边界和 goroutine 携带截止时间、取消信号和其他请求范围的值。
context.Context机器为每个请求创建一个net/http，并可使用该Context()方法。
等待几秒钟后再向客户端发送回复。这可以模拟服务器正在进行的一些工作。
在工作期间，请密切关注上下文的 Done()通道，以查看是否收到我们应该取消工作并尽快返回的信号。
上下文的Err()方法返回一个错误，解释为什么Done()通道被关闭。
和以前一样，我们在“/hello”路线上注册我们的处理程序，并开始服务。

在后台运行服务器。
模拟客户端请求/hello，在启动后不久按 Ctrl+C 来发出取消信号
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
*
您提供的Go代码实现了一个简单的HTTP服务器，该服务器在访问/hello路径时会启动hello处理器函数。
这个处理器函数使用select语句来等待两个事件中的一个：要么是10秒钟的超时事件，要么是请求上下文（ctx）被取消的事件。
下面是代码的详细解释以及它的工作原理：
*/
/**
// hello 函数是一个HTTP处理器，当访问/hello路径时会被调用。
// 它首先打印一条消息表示处理器开始工作，然后使用一个select语句来等待。
// select语句有两个case：一个是等待10秒钟的超时事件，另一个是等待请求上下文被取消的事件。
*/
func hello1(w http.ResponseWriter, req *http.Request) {
	// 获取请求的上下文
	ctx := req.Context()
	fmt.Println("server: hello1 handler started")
	// 使用defer确保在函数结束时打印结束消息
	defer fmt.Println("server: hello1 handler ended")
	// 使用select语句来等待两个事件中的一个
	select {
	case <-time.After(10 * time.Second):
		// 如果10秒钟超时事件先到，就向HTTP响应中写入"hello\n"字符串
		fmt.Fprintf(w, "hello1\n")
	case <-ctx.Done():
		// 如果请求上下文被取消的事件先到，就获取错误并打印，然后向客户端返回500内部服务器错误
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

// main 函数是程序的入口点。
// 它注册了/hello路径的处理器为hello函数，然后启动HTTP服务器监听本地的8090端口。
func main() {

	http.HandleFunc("/hello1", hello1)
	http.ListenAndServe(":8090", nil)
}

/**
当您运行这个程序时，它会在本地的8090端口上启动一个HTTP服务器。如果您访问http://localhost:8090/hello，服务器会启动hello处理器函数。

处理器函数首先会打印"server: hello handler started"来表示它开始工作了。
然后，它会进入一个select语句，等待两个事件中的一个发生：
如果10秒钟内没有其他事情发生（即没有上下文取消的事件），那么它会向HTTP响应中写入"hello\n"字符串，并且最终打印"server: hello handler ended"来表示处理器函数结束了。
如果在这10秒钟内请求上下文被取消了（可能是因为客户端关闭了连接或者服务器内部发生了某种超时或取消操作），那么它会捕获到ctx.Done()事件，获取错误并打印，然后向客户端返回一个500内部服务器错误。
这个处理器函数展示了如何在Go中使用上下文（context.Context）来管理请求的生命周期，以及如何使用select语句来等待多个可能的事件。在实际应用中，这种模式可以用于实现超时、取消操作或其他需要等待多个异步事件完成的场景。
*/
