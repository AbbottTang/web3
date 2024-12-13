// 关闭通道表示不再向其发送任何值。这对于向通道的接收者传达完成信息很有用。
// 在这个例子中，我们将使用一个jobs通道将要完成的工作从main()goroutine 传达给工作 goroutine。
// 当我们没有更多工作需要工作时，我们将使用close通道jobs。
// 这是工作 goroutine。它反复从 接收jobs。j, more := <-jobs在这个特殊的 2 值形式的接收中，
// 如果已被d 并且通道中的所有值都已被接收，more则值将是。
// 我们使用它来通知我们何时完成了所有工作。falsejobsclosedone
// 这将通过通道向工作者发送 3 个作业jobs ，然后关闭它。
// 我们使用之前看到的同步方法等待工作者 。
// 从已关闭的通道读取数据会立即成功，并返回基础类型的零值。
// 第二个可选返回值是true接收的值是否由成功发送到通道的发送操作传送，
// 或者false是否是由于通道已关闭且为空而生成的零值。
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	//在 Go 语言中，j, more := <-jobs 是一个从通道 jobs 接收值的语句，同时它还检查通道是否已经关闭。
	//这里的 jobs 是一个可以传递数据的通道（channel），而 j 是用来接收从通道中传递过来的值的变量，
	//more 是一个布尔类型的变量，用来指示通道是否已经关闭。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
