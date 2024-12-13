package main

import "fmt"

// 通道上的基本发送和接收都是阻塞的。
// 但是，我们可以使用selectwithdefault子句来实现非阻塞发送、接收，甚至非阻塞多路select发送。
// 这是一个非阻塞接收。
// 如果上有一个值可用messages，则将使用该值select获取<-messages case。如果没有，它将立即获取该default案例。
// 非阻塞发送的工作原理类似。这里msg 无法发送到messages通道，因为通道没有缓冲区，也没有接收器。
// 因此default选择了这种情况。
// 我们可以case在子句上方使用多个 sdefault 来实现多路非阻塞选择。
// 这里我们尝试在messages和 上进行非阻塞接收signals。
func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//go func() {
	//	messages <- "hello"
	//}()
	go func() {
		signals <- true
	}()
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
