/*
*

环境变量是向 Unix 程序传递配置信息 的通用机制。让我们看看如何设置、获取和列出环境变量。
要设置键/值对，请使用os.Setenv。要获取键的值，请使用os.Getenv。如果环境中不存在该键，这将返回一个空字符串。
用于os.Environ列出环境中的所有键/值对。这将返回格式为 的字符串片段KEY=value。
您可以使用strings.SplitN它们来获取键和值。这里我们打印所有键。\

运行程序表明，我们获取了FOO程序中设置的值，但该值 BAR是空的。
环境中的密钥列表取决于您的特定机器。
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 设置环境变量FOO的值为"1"
	os.Setenv("FOO", "1")
	// 打印环境变量FOO的值
	fmt.Println("FOO:", os.Getenv("FOO"))
	// 尝试打印环境变量BAR的值，但BAR可能并未设置
	fmt.Println("BAR:", os.Getenv("BAR"))
	// 打印一个空行作为分隔
	fmt.Println()
	// 遍历所有环境变量
	for _, e := range os.Environ() {
		// 使用strings.SplitN将环境变量字符串按"="分割成两部分，最多分割成2个部分
		pair := strings.SplitN(e, "=", 2)
		// 打印环境变量的名称（即分割后的第一部分）
		fmt.Println(pair[0])
	}
}
