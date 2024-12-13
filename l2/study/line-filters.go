package main

//bufio：用于缓冲I/O，这里主要用于读取输入。
//fmt：用于格式化I/O，这里主要用于输出文本。
//os：提供了操作系统功能接口，这里用于获取标准输入和输出。
//strings：提供了字符串操作功能，这里用于将文本转换为大写。
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//创建一个bufio.Scanner实例scanner，用于从os.Stdin（标准输入）读取数据。
	//bufio.NewScanner(os.Stdin)会创建一个新的Scanner，它以标准输入作为数据源。
	scanner := bufio.NewScanner(os.Stdin)
	//使用for scanner.Scan()循环读取输入，直到没有更多的输入（例如，用户输入EOF信号，通常是Ctrl+D（Unix/Linux/Mac）或Ctrl+Z（Windows））。
	//在循环体内，scanner.Text()获取当前读取的行文本，strings.ToUpper将其转换为大写。
	//使用fmt.Println(ucl)将转换后的大写文本输出到标准输出。
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	//循环结束后，检查scanner.Err()是否返回错误。如果有错误，err将不为nil。
	//如果存在错误，使用fmt.Fprintln(os.Stderr, "error:", err)将错误信息输出到标准错误输出（通常是屏幕），
	//并通过os.Exit(1)退出程序，1表示程序因为错误而退出。
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
