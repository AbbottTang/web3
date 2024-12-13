// goroutine是一个轻量级的执行线程。
package main

import (
	"fmt"
	"time"
)

func f3(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// 要在 goroutine 中调用此函数，请使用 go f(s)。这个新的 goroutine 将与调用 goroutine 同时执行。
func main() {

	f3("direct")

	go f3("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
