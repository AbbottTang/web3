package main

import (
	"fmt"
	"os"
)

// 命令行参数 是参数化程序执行的常用方法。例如，go run hello.go使用run和 hello.go参数go。
// os.Args提供对原始命令行参数的访问。请注意，此切片中的第一个值是程序的路径，并os.Args[1:] 保存程序的参数。
func main() {

	//使用了os.Args来获取命令行参数。os.Args是一个字符串切片，其中包含了命令行参数的列表，
	//第一个元素（索引为0）通常是程序的名称，随后的元素则是传递给程序的参数。

	// 包含程序名称的所有命令行参数
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:] // 不包含程序名称的命令行参数

	arg := os.Args[3] // 尝试获取第四个命令行参数

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
