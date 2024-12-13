package main

import "os"

//通常表示发生了意外错误。我们主要用它panic来快速失败，
//以应对在正常操作期间不应发生的错误，或者我们无法妥善处理的错误。
//我们将在整个站点中使用 panic 来检查意外错误。这是站点上唯一设计为 panic 的程序。
//panic 的一个常见用途是，如果函数返回一个我们不知道如何处理（或不想处理）的错误值，则中止。
//以下是 panicking 的示例，如果我们在创建新文件时遇到意外错误。

//运行该程序将导致其崩溃，打印错误消息和 goroutine 跟踪，并以非零状态退出。
//当第一次 panic inmain触发时，程序会退出而不执行其余代码。
//如果您想看到程序尝试创建临时文件，请注释掉第一次 panic out

func main() {

	//panic("a problem")

	_, err := os.Create("C:/Users/Administrator/Desktop/tmp/test.txt")
	if err != nil {
		panic(err)
	}
}
