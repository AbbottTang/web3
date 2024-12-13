package main

import "fmt"

// 在前面的例子中，我们看到了如何for对 range基本数据结构进行迭代。我们也可以使用此语法对从通道接收的值进行迭代。
// 我们将迭代queue通道中的 2 个值。
// 这range将对从 接收到的每个元素进行迭代queue。由于我们close已使用上面的通道，因此迭代在接收到 2 个元素后终止。
// 这个例子还表明，可以关闭非空通道，但仍然可以接收剩余的值。
func main() {
	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	queue <- "two"
	queue <- "two"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}
