// 通道channel是连接并发 goroutine 的管道。
// 你可以将值从一个 goroutine 发送到通道，并在另一个 goroutine 中接收到这些值。
package main

import "fmt"

func main() {
	//使用 创建一个新通道make(chan val-type)。通道按其传达的值进行分类。
	messages := make(chan string)
	//使用语法将值发送到通道channel <- 。这里我们从一个新的 goroutine 发送"ping" 到messages 上面创建的通道。
	go func() { messages <- "ping" }()
	//该<-channel语法从通道接收"ping"一个值。在这里我们将接收上面发送的消息并将其打印出来
	msg := <-messages
	fmt.Println(msg)
	//当我们运行程序时，"ping"消息通过我们的通道成功地从一个 goroutine 传递到另一个。
	//默认情况下，发送和接收会阻塞，直到发送者和接收者都准备好为止。
	//此属性允许我们在程序结束时等待消息，"ping" 而无需使用任何其他同步
}
