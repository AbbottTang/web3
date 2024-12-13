package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

//在上一个示例中，我们使用了显式锁定和 互斥锁来同步多个 goroutine 之间对共享状态的访问。
//另一种选择是使用 goroutine 和通道的内置同步功能来实现相同的结果。
//这种基于通道的方法符合 Go 的共享内存理念，即通过通信并让每段数据都归 1 个 goroutine 所有。
//在此示例中，我们的状态将由单个 goroutine 拥有。这将保证数据永远不会因并发访问而损坏。
//为了读取或写入该状态，其他 goroutine 将向拥有该状态的 goroutine 发送消息并接收相应的回复。
//这些readOp和writeOp struct封装了这些请求以及拥有该状态的 goroutine 的响应方式。
//与之前一样，我们将计算执行了多少个操作
//reads和通道writes将分别被其他 goroutines 用于发出读写请求。
//这是拥有 的 goroutine state，它与上例一样是一个映射，但现在是状态 goroutine 的私有映射。
//此 goroutine 反复选择reads和writes通道，并在请求到达时做出响应。
//响应的执行方式是，首先执行请求的操作，然后在响应通道上发送一个值resp来指示成功（在 的情况下为所需的值reads）。
//这将启动 100 个 goroutine，通过通道向拥有状态的 goroutine 发出读取操作reads。
//每次读取都需要构造一个readOp，通过通道发送reads，然后通过提供的通道接收结果resp。

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64
	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()
	for r := 0; r < 100; r++ {
		go func() {
			//for {
			read := readOp{
				key:  rand.Intn(5),
				resp: make(chan int),
			}
			reads <- read
			<-read.resp
			atomic.AddUint64(&readOps, 1)
			time.Sleep(time.Millisecond)

			//}
		}()
	}
	for w := 0; w < 10; w++ {
		go func() {
			//for {
			write := writeOp{
				key:  rand.Intn(5),
				val:  rand.Intn(100),
				resp: make(chan bool),
			}
			writes <- write
			<-write.resp
			atomic.AddUint64(&writeOps, 1)
			time.Sleep(time.Millisecond)
			//}
		}()
	}
	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
