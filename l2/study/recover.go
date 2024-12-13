package main

import "fmt"

//Go 可以使用内置函数从 panic 中恢复。A可以阻止 a中止程序，而是让其继续执行。recoverrecoverpanic
//举个例子，这个功能很有用：如果某个客户端连接出现严重错误，服务器不会崩溃。
//相反，服务器会关闭该连接并继续为其他客户端提供服务。
//事实上，这正是 Gonet/http 默认为 HTTP 服务器所做的。
//recover必须在延迟函数内调用。当封闭函数发生恐慌时，延迟将激活，并且recover其中的调用将捕获恐慌。
//的返回值recover是调用时引发的错误panic。
//由于出现恐慌，此代码将无法运行mayPanic。执行将main在恐慌发生时停止，并在延迟闭包中恢复。

func mayPanic() {
	panic("a problem")
}

func main() {

	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
