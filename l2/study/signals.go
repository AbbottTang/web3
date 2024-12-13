package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
*
有时我们希望 Go 程序能够智能地处理Unix 信号。例如，我们可能希望服务器在收到 时正常关闭SIGTERM，或者命令行工具在收到 时停止处理输入SIGINT。以下是如何在 Go 中使用通道处理信号。
Go 信号通知通过os.Signal 在通道上发送值来工作。我们将创建一个通道来接收这些通知。请注意，此通道应该是缓冲的。
signal.Notify注册给定的通道来接收指定信号的通知。
我们可以从主函数中接收sigs，但让我们看看如何在单独的 goroutine 中完成此操作，以演示更现实的正常关闭场景。
这个 goroutine 执行一个阻塞信号接收。当它收到一个信号时，它会将其打印出来，然后通知程序它可以完成了。
程序会在这里等待，直到得到预期的信号（如上面的 goroutine 发送值所示done）然后退出。

当我们运行此程序时，它将阻塞以等待信号。通过输入ctrl-C（终端显示为^C），我们可以发送SIGINT信号，导致程序打印interrupt然后退出
*/
func main() {
	// 创建一个通道来接收操作系统信号

	sigs := make(chan os.Signal, 1)
	// 告诉Go的signal包，当接收到SIGINT或SIGTERM信号时，将它们发送到sigs通道

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// 创建一个布尔通道，用于同步goroutine的结束

	done := make(chan bool, 1)
	// 启动一个新的goroutine来处理接收到的信号

	go func() {

		sig := <-sigs    // 从sigs通道接收信号
		fmt.Println()    // 打印一个空行
		fmt.Println(sig) // 打印接收到的信号
		done <- true     // 向done通道发送true，表示goroutine已完成
	}()

	fmt.Println("awaiting signal") // 打印等待信号的提示
	<-done                         // 等待goroutine完成处理（即等待从done通道接收到true）
	fmt.Println("exiting")         // 打印退出程序的提示
	/**
	‌在本地终端/命令行中‌：当你在本地终端或命令行中运行这个程序时，你可以通过按Ctrl+C来发送SIGINT信号，或者通过其他方式（如进程管理工具）发送SIGTERM信号。程序会捕获这些信号，打印出它们，然后退出。

	‌在Docker容器或某些其他隔离环境中‌：如果你将这个程序部署到Docker容器或其他隔离环境中，发送信号的方式可能会有所不同。例如，在Docker中，你可能需要使用docker kill命令来发送SIGTERM信号。

	‌在Windows上‌：虽然SIGINT和SIGTERM是Unix/Linux特有的信号，但Go的os/signal包在Windows上也提供了对这些信号的支持（通过模拟的方式）。在Windows上按Ctrl+C通常会发送一个类似于SIGINT的信号。

	注意事项
	确保你的Go环境已经正确设置，并且你使用的是支持并发编程的Go版本。
	如果你在Windows上运行这个程序，并且发现信号没有被正确处理，请检查你的Go版本是否是最新的，因为旧版本的Go可能在Windows上的信号处理方面存在限制。
	在生产环境中，使用信号来优雅地关闭程序是一种常见的做法，因为它允许程序在退出之前完成必要的清理工作（如关闭数据库连接、保存状态等）。
	*/
}
