package main

import (
	"flag"
	"fmt"
	"os"
)

/*
*

一些命令行工具，如go工具或git 有许多子命令，每个子命令都有自己的一组标志。
例如，go build和go get是go工具的两个不同子命令。该flag包让我们可以轻松定义具有自己标志的简单子命令。
程序支持两个子命令：foo 和 bar。每个子命令都有自己的命令行参数，并且程序会根据用户输入的第一个参数（即子命令名称）来决定解析哪个子命令的参数。
这是通过使用flag.NewFlagSet来为每个子命令创建一个独立的标志集实现的。
*/
func main() {
	// 为'foo'子命令创建一个标志集，并设置当解析错误时退出程序
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable the command")
	fooName := fooCmd.String("name", "", "the name of the command")

	// 为'bar'子命令创建一个标志集，并设置当解析错误时退出程序
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "the level of the command")

	// 检查是否有至少一个参数（子命令名称）
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
	// 根据第一个参数（子命令名称）来解析相应的子命令参数
	switch os.Args[1] {
	case "foo":
		// 解析'foo'子命令的参数（跳过程序名称和子命令名称）
		fooCmd.Parse(os.Args[2:])
		// 打印'foo'子命令的参数和剩余的非标志参数
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		// 解析'bar'子命令的参数（跳过程序名称和子命令名称）
		barCmd.Parse(os.Args[2:])
		// 打印'bar'子命令的参数和剩余的非标志参数
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		// 如果输入的子命令名称不是'foo'或'bar'，则打印错误信息并退出程序
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
	//./command-line-subcommands foo -enable -name=joe a1 a2
	//./command-line-subcommands bar -level 8 a1
	// ./command-line-subcommands bar -enable a1

}
