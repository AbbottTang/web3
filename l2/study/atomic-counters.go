package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Go 中管理状态的主要机制是通过通道进行通信。我们在工作池中看到了这一点。不过，还有其他一些管理状态的选项。
//在这里，我们将研究如何使用sync/atomic包来访问多个 goroutine 的原子计数器。
//我们将使用原子整数类型来表示我们的（始终为正的）计数器。
//WaitGroup 将帮助我们等待所有 goroutine 完成其工作。
//我们将启动 50 个 goroutine，每个 goroutine 将计数器增加 1000 次。
//为了原子地增加计数器，我们使用Add/
///等到所有 goroutine 都完成
//这里没有 goroutines 正在写入“ops”，
//但使用 Load它来原子地读取一个值是安全的，即使其他 goroutines 正在（原子地）更新它。
//
//我们预计会得到 50,000 次操作。如果我们使用非原子整数并用 增加它 ops++，
//我们可能会得到不同的数字，在运行之间发生变化，
//因为 goroutine 会相互干扰。此外，在使用该标志运行时，我们会遇到数据争用失败 -race。

func main() {
	var ops atomic.Uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops.Load())
}
