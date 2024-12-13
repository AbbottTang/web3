package main

import (
	"fmt"
	"time"
)

// 速率限制 是控制资源利用率和维护服务质量的重要机制。Go 优雅地使用 goroutines、channels 和tickers支持速率限制。
// 首先，我们来了解一下基本速率限制。假设我们想限制对传入请求的处理。我们将通过同名的通道来处理这些请求
// 此limiter通道每 200 毫秒将接收一个值。这是我们的速率限制方案中的调节器。/
// 通过在处理每个请求之前阻止从通道接收limiter，我们将自己限制为每 200 毫秒 1 个请求。//
// 我们可能希望在速率限制方案中允许短时间突发请求，同时保持总体速率限制。我们可以通过缓冲限制器通道来实现这一点。此burstyLimiter 通道将允许最多 3 个事件的突发。
// 填满通道以表示允许突发。
// 每 200 毫秒我们将尝试向中添加一个新值burstyLimiter，最多为 3。/
// 现在模拟另外 5 个传入请求。其中前 3 个将受益于 的突发功能burstyLimiter。
// 运行我们的程序，我们看到第一批请求每~200毫秒处理一次，正如预期的那样。
// 对于第二批请求，由于突发速率限制，我们会立即处理前 3 个请求，然后以每个请求约 200 毫秒的延迟处理剩下的 2 个请求

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(1000 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
