package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	//因为这个通道是缓冲的，所以我们可以把这些值发送到通道中，而不需要相应的并发接收
	done <- true
}

func main() {
	//使用 创建一个新通道make(chan val-type)。通道按其传达的值进行分类。
	//这里我们有make一个bool通道，最多可缓冲 1 个值
	done := make(chan bool, 1)
	// 要在 goroutine 中调用此函数，请使用 go f(s)。这个新的 goroutine 将与调用 goroutine 同时执行。
	go worker(done)

	msg := <-done
	fmt.Println(msg)
}
