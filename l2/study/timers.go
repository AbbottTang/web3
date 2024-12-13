package main

import (
	"fmt"
	"time"
)

//我们经常希望在未来的某个时间点执行 Go 代码，或者以某个间隔重复执行。
//Go 的内置 计时器和代码行距功能使这两项任务都变得容易。
//我们首先了解计时器，然后了解代码行距。
//计时器代表未来的单个事件。
//您告诉计时器要等待多长时间，它会提供一个将在那时收到通知的通道。此计时器将等待 2 秒。
//<-timer1.C计时器的通道被阻塞直到C 它发送一个指示计时器已触发的值。
//如果您只是想等待，您可以使用 time.Sleep。计时器有用的一个原因是您可以在计时器触发之前取消它。以下是一个例子。
//给予timer2足够的时间使其点火（如果它要点火的话），以表明它实际上已经停止了。
//第一个计时器将在我们启动程序后约 2 秒内触发，但第二个计时器应该在触发之前停止。

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	msg := <-timer1.C
	fmt.Println(msg)
	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		msg2 := <-timer2.C
		fmt.Println(msg2)
		fmt.Println("timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
	//go func() 启动了一个新的 goroutine 来执行匿名函数。
	//由于 goroutine 是异步执行的，主程序会继续执行 time.Sleep 语句，而不会等待 goroutine 完成。
	//我们在这里使用 time.Sleep 是为了给 goroutine 提供一个执行的机会，否则主程序可能会立即退出，
	//导致 goroutine 中的代码没有机会被执行。
	time.Sleep(2 * time.Second)
}

/**
在 Go 语言中，go func() 用于启动一个新的 goroutine，该 goroutine 会并发地执行与之关联的函数。然而，goroutine 的执行是异步的，这意味着它不会阻塞主 goroutine（即主线程）的执行。因此，如果主 goroutine 在启动新的 goroutine 后立即结束，那么整个程序可能会在主 goroutine 完成时终止，而新启动的 goroutine 可能还没有机会执行或只执行了部分代码。

使用 time.Sleep 是一种简单的方法来确保主 goroutine 在给新启动的 goroutine 足够时间执行之前不会结束。这是一种人为的延迟，用于模拟或等待某些操作完成。然而，它并不是一种可靠的同步机制，因为：

你无法准确知道 goroutine 需要多长时间来完成其任务。
如果 goroutine 执行的时间比 time.Sleep 指定的时间更长，那么主 goroutine 仍然可能会在 goroutine 完成之前结束。
如果 goroutine 执行的时间比 time.Sleep 指定的时间更短，那么主 goroutine 将会不必要地等待。
因此，time.Sleep 通常只用于示例代码、测试或调试目的，而不应该用于生产代码中的同步。

在生产代码中，你应该使用更可靠的同步机制，如通道（channel）、sync.WaitGroup 或其他并发原语来确保 goroutine 之间的正确协调和同步。这些机制允许你以更安全、更可预测的方式等待 goroutine 完成其任务，而不会依赖于不确定的延迟
*/
