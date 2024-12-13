package main

import "fmt"

/*
*

默认情况下，通道是无缓冲的chan <-，这意味着，只有当有相应的接收 ( <- chan) 准备好接收发送的值时，
它们才会接受发送 ( )。缓冲通道接受有限数量的值，而这些值没有相应的接收器。
*/
func main() {
	//这里我们有make一个字符串通道，最多可缓冲 2 个值
	messages := make(chan string, 2)
	//因为这个通道是缓冲的，所以我们可以把这些值发送到通道中，而不需要相应的并发接收
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
