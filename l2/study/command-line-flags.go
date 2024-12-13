package main

import (
	"flag"
	"fmt"
)

//命令行标志

func main() {
	// 定义一个指向string类型的指针，并设置默认值、名称和帮助信息
	wordPtr := flag.String("word", "foo", "a string")

	// 定义一个指向int类型的指针，并设置默认值、名称和帮助信息
	numbPtr := flag.Int("numb", 42, "an int")

	// 定义一个指向bool类型的指针，并设置默认值、名称和帮助信息
	forkPtr := flag.Bool("fork", false, "a bool")

	// 定义一个string类型的变量，并使用StringVar方法为其设置默认值、名称和帮助信息
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 解析命令行参数
	flag.Parse()

	// 打印解析后的参数值
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)

	// 打印剩余的命令行参数（非flag参数）
	fmt.Println("tail:", flag.Args())

	//go run yourprogram.go -word=hello -numb=123 -fork -svar=world extraArg1 extraArg2
	//word: hello
	//numb: 123
	//fork: true
	//svar: world
	//tail: [extraArg1 extraArg2]

	//-word=hello 覆盖了word参数的默认值，所以输出是hello。
	//-numb=123 覆盖了numb参数的默认值，所以输出是123。
	//-fork 没有指定值，因为fork是bool类型，所以它的值被设置为true。
	//-svar=world 覆盖了svar参数的默认值，所以输出是world。
	//extraArg1 extraArg2 是非flag参数，它们被flag.Args()捕获并打印出来。
	//如果您没有提供任何命令行参数，程序将使用默认值打印出参数值，并且flag.Args()将返回一个空切片。

	//使用-h或--help标志来获取命令行程序的自动生成的帮助文本。
	//./command-line-flags -h
	//如果您提供未在包中指定的标志 flag，程序将打印错误消息并再次显示帮助文本。
	//./command-line-flags -wat
}
