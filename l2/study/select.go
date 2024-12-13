package main

import (
	"fmt"
	"time"
)

// 在我们的示例中，我们将跨两个渠道进行选择。
// 我们将select同时等待这两个值，并在每个值到达时打印它们。
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
	//每个通道将在一定时间后收到一个值，以模拟在并发 goroutine 中执行的阻塞 RPC 操作。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
