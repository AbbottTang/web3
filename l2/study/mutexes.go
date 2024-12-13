package main

import (
	"fmt"
	"sync"
)

//在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器状态。
//对于更复杂的状态，我们可以使用互斥锁 来安全地跨多个 goroutine 访问数据。
//容器保存计数器映射；由于我们想从多个 goroutine 同时更新它，因此我们添加了一个Mutex来同步访问。
//请注意，互斥锁不能被复制，因此如果 struct传递它，应该通过指针来完成。
//访问之前锁定互斥锁counters；在函数末尾使用defer 语句将其解锁。
//请注意，互斥锁的零值可以按原样使用，因此这里不需要初始化。
//此函数循环增加一个命名的计数器。
//同时运行多个 goroutine；注意它们都访问同一个Container，其中两个访问同一个计数器。
//等待 goroutines 完成
//运行程序表明计数器按预期更新。
/**
在您的Go语言代码示例中，您定义了一个Container结构体，该结构体包含一个sync.Mutex类型的锁mu和一个字符串到整数的映射counters。您还定义了一个inc方法，用于安全地增加指定计数器的值。在main函数中，您使用了sync.WaitGroup来同步三个goroutine，这些goroutine并发地调用inc方法来增加计数器的值。

这里是代码的详细解释：

Container 结构体
Container结构体包含两个字段：
mu：一个sync.Mutex类型的锁，用于确保对counters映射的并发访问是安全的。
counters：一个字符串到整数的映射，存储了不同名称的计数器及其当前值。
inc 方法
inc方法接收一个name字符串参数，表示要增加的计数器的名称。
方法内部首先调用c.mu.Lock()来加锁，确保对counters的访问是线程安全的。
使用defer c.mu.Unlock()来确保在方法返回时释放锁，无论方法是否正常结束。
然后，通过c.counters[name]++来增加指定计数器的值。
main 函数
在main函数中，首先创建了一个Container实例c，并初始化counters映射，其中包含两个计数器"a"和"b"，它们的初始值都是0。
声明了一个sync.WaitGroup类型的变量wg，用于等待goroutine的完成。
定义了一个doIncrement函数，该函数接收一个计数器名称name和一个整数n，表示要增加计数器的次数。函数内部使用一个for循环来多次调用c.inc(name)方法。
使用wg.Add(3)来增加WaitGroup的计数器，表示有三个goroutine需要等待。
启动了三个goroutine，它们分别调用doIncrement函数来增加"a"计数器和"b"计数器的值。其中，有两个goroutine增加"a"计数器的值，一个goroutine增加"b"计数器的值。
调用wg.Wait()来阻塞main函数，直到所有的goroutine都完成。
最后，打印出counters映射的值，显示每个计数器的最终值。
代码执行流程
程序启动，执行main函数。
创建Container实例c，并初始化counters映射。
声明WaitGroup变量wg，并增加其计数器到3。
启动三个goroutine，它们并发地调用doIncrement函数来增加计数器的值。
wg.Wait()阻塞main函数，直到所有的goroutine都调用wg.Done()来减少WaitGroup的计数器。
所有的goroutine完成后，WaitGroup的计数器变为零，wg.Wait()解除阻塞。
打印出counters映射的值，显示每个计数器的最终值。
main函数结束，程序退出。
潜在的问题和改进
在这个示例中，代码是线程安全的，因为对counters映射的访问被mu锁保护。然而，有一些小的改进点：

在doIncrement函数中调用wg.Done()时，最好使用defer来确保在函数结束时调用，即使在循环中发生panic也能保证wg.Done()被调用。但在您的代码中，由于循环体内没有可能引发panic的操作，所以这不是必需的。
如果counters映射的键是动态生成的，并且数量很大，您可能需要考虑使用sync.Map或其他并发安全的映射来实现，以避免锁的争用和性能瓶颈。但在您的示例中，键是固定的，所以使用普通的映射和锁是合适的。
*/
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}
func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)
	go doIncrement("a", 100)
	go doIncrement("a", 500)
	go doIncrement("b", 100)
	wg.Wait()
	fmt.Println(c.counters)
}
