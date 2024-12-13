package main

import (
	"fmt"
	"time"
)

// 对于连接到外部资源或需要限制执行时间的程序来说，超时select非常重要。借助通道和，在 Go 中实现超时既简单又优雅。
// 在我们的示例中，假设我们正在执行一个外部调用，该调用在 2 秒后在通道上返回其结果c1 。
// 请注意，该通道是缓冲的，因此 goroutine 中的发送是非阻塞的。
// 这是一种常见的模式，用于在通道从未被读取的情况下防止 goroutine 泄漏
// 以下是select超时的实现。 res := <-c1等待结果，并<-time.After 等待在 1 秒的超时后发送值。
// 由于select会继续执行第一个已准备好的接收，因此如果操作所用时间超过允许的 1 秒，我们将采用超时情况。
func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
