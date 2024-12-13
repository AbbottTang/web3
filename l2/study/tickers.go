package main

import (
	"fmt"
	"time"
)

// 计时器适用于您想要在未来做某事一次的情况 -滴答器适用于您想要定期重复做某事的情况。
// 以下是滴答器的示例，它会定期滴答，直到我们将其停止。
// 行情指示器使用与计时器类似的机制：一个用于发送值的通道。
// 在这里，我们将使用 select通道上的内置函数来等待每 500 毫秒到达的值
// 行情指示器可以像计时器一样停止。行情指示器停止后，其通道上将不再接收任何值。我们将在 1600 毫秒后停止行情指示器。
// 当我们运行这个程序时，在我们停止它之前，它会滴答作响 3 次。
func main() {
	// 创建一个每 500 毫秒触发一次的 Ticker
	ticker := time.NewTicker(time.Millisecond * 500)
	// 创建一个布尔类型的通道，用于通知 goroutine 停止
	done := make(chan bool)
	// 启动一个新的 goroutine 来处理 Ticker 的触发事件
	go func() {
		for {
			select {
			// 当接收到 done 信号时，退出循环并返回
			case <-done:
				return
				// 当 Ticker 触发时，打印当前时间
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	// 主 goroutine 等待 1600 毫秒
	time.Sleep(time.Millisecond * 1600)
	// 停止 Ticker，防止它继续触发
	ticker.Stop()
	// 向 done 通道发送一个信号，通知另一个 goroutine 停止
	done <- true
	// 打印一条消息，表示 Ticker 已停止
	fmt.Println("Ticker stopped")
}

/**
time.NewTicker(500 * time.Millisecond) 创建一个新的 time.Ticker，它每 500 毫秒向其内部的 C 通道发送一个当前时间的值。

done := make(chan bool) 创建一个布尔类型的通道 done，用于在需要时通知 goroutine 停止执行。

使用 go func() 启动一个新的 goroutine，在这个 goroutine 中包含一个无限循环 for，该循环使用 select 语句来监听两个通道：done 和 ticker.C。

在主 goroutine 中，使用 time.Sleep(1600 * time.Millisecond) 让主程序等待 1600 毫秒。

1600 毫秒后，主 goroutine 调用 ticker.Stop() 停止 Ticker，防止它继续向 ticker.C 发送时间值。

然后，主 goroutine 向 done 通道发送一个 true 值，通知另一个 goroutine 退出循环并返回。

最后，主 goroutine 打印一条消息 "Ticker stopped"，表示 Ticker 已停止。

在这个例子中，由于 Ticker 每 500 毫秒触发一次，而主程序等待了 1600 毫秒，因此在 Ticker 被停止之前，它大约会触发三次（分别在 500ms、1000ms 和 1500ms 时）。之后，Ticker 被停止，并且 done 通道通知另一个 goroutine 停止执行。
*/
