package main

import (
	"fmt"
	"time"
)

//在这个例子中，我们将研究如何使用 goroutines 和通道实现工作池。
//这是工作器，我们将运行它的多个并发实例。这些工作器将在jobs通道上接收工作，并在上发送相应的结果results。
//我们将每个作业休眠一秒钟以模拟一项昂贵的任务。
//为了使用我们的员工队伍，我们需要向他们发送工作并收集他们的成果。我们为此建立了 2 个渠道
//这将启动 3 个工人，最初被阻止是因为目前还没有工作。
//这里我们发送 5 jobs，然后发送close该频道来表明这就是我们所有的工作。
//最后，我们收集所有工作结果。这也确保工作 goroutine 已完成。等待多个 goroutine 的另一种方法是使用WaitGroup。

// 我们正在运行的程序显示了由各个 worker 执行的 5 个作业。尽管总共执行了大约 5 秒钟的工作，但该程序仅花费了大约 2 秒钟，因为有 3 个 worker 同时运行。

// jobs <-chan int：这是一个只接收类型的通道参数。
// result chan<- int：这是一个只发送类型的通道参数。
func worker1(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
}
func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for w := 1; w <= 3; w++ {
		go worker1(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
