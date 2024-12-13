package main

import (
	"bufio"
	"fmt"
	"net/http"
)

/*
*
Go 标准库对net/http 软件包中的 HTTP 客户端和服务器提供了出色的支持。
在此示例中，我们将使用它发出简单的 HTTP 请求。
向服务器发出 HTTP GET 请求。是创建 对象并调用其方法的http.Get便捷方式；
它使用 具有有用的默认设置的对象。http.ClientGethttp.DefaultClient
*/
func main() {

	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 100; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
