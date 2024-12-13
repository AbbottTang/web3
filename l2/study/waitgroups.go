package main

import (
	"fmt"
	"sync"
	"time"
)

// 要等待多个 goroutine 完成，我们可以使用等待组。
// 这是我们将在每个 goroutine 中运行的函数。
// 睡眠来模拟一项昂贵的任务。
// 这个 WaitGroup 用于等待这里启动的所有 goroutine 完成。
// 注意：如果将 WaitGroup 明确传递给函数，则应通过指针来完成。
// 启动多个 goroutine 并增加每个 goroutine 的 WaitGroup 计数器。
// /将 worker 调用包装在一个闭包中，以确保告知 WaitGroup 该 worker 已完成。
// 这样，worker 本身就不必了解其执行中涉及的并发原语
// /阻塞直到 WaitGroup 计数器返回到 0；所有工人都会收到通知说他们已经完成工作。
// 请注意，这种方法没有直接的方式来传播来自工作器的错误。对于更高级的用例，请考虑使用 errgroup 包。
// /每次调用时，工作者启动和完成的顺序可能会有所不同。
func worker3(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func main() {
	//‌声明WaitGroup‌：var wg sync.WaitGroup声明了一个sync.WaitGroup类型的变量wg。
	//WaitGroup用于等待一组goroutine完成。
	//‌启动worker‌：使用一个for循环启动了5个worker goroutine。在每个goroutine中，
	//首先通过wg.Add(1)增加了WaitGroup的计数器，表示有一个新的goroutine需要等待。
	//然后，使用go func()启动了一个匿名函数作为goroutine，该函数内部调用了worker函数，
	//并在最后通过defer wg.Done()在函数返回时减少了WaitGroup的计数器。
	//这里需要注意的是，由于i变量在for循环中被捕获，并且goroutine的执行是并发的，
	//因此如果直接在goroutine内部使用i，可能会导致数据竞争或不可预测的行为。
	//然而，在这个特定的示例中，由于i只是用作worker函数的参数，并且worker函数只是简单地打印了它，
	//所以没有出现问题。但在更复杂的场景中，您可能需要使用局部变量或同步机制来确保数据的正确性。
	//‌等待worker完成‌：wg.Wait()阻塞了main函数，直到WaitGroup的计数器变为零，表示所有的worker都已经完成。
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker3(i)
		}()
	}
	wg.Wait()
}
