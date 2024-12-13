package main

import (
	"fmt"
	"os/exec"
)

/*
*
生成进程
我们将从一个简单的命令开始，该命令不接受任何参数或输入，只会将一些内容打印到标准输出。
exec.Command助手会创建一个对象来表示这个外部进程。
该Output方法运行命令，等待其完成并收集其标准输出。如果没有错误，dateOut将保存包含日期信息的字节
Output如果执行命令时出现问题（例如路径错误），
或者 命令运行但以非零返回代码退出，则其他方法Command将返回 。*exec.Error*exec.ExitError
接下来我们来看一个稍微复杂一点的案例，我们将数据通过管道传输到其外部进程 stdin并从其收集结果stdout。
在这里，我们明确地抓住输入/输出管道，启动进程，向其中写入一些输入，读取结果输出，最后等待进程退出

我们在上面的例子中省略了错误检查，但您可以if err != nil对所有错误检查使用通常的模式。
我们也只收集结果，但您可以以完全相同的方式StdoutPipe 收集。StderrPipe

请注意，生成命令时，我们需要提供明确划定的命令和参数数组，而不是只传入一个命令行字符串。
如果要生成带有字符串的完整命令，可以使用bash的-c 选项：

生成的程序返回的输出与我们直接从命令行运行它们的输出相同。
go run spawning-processes.go
日期没有-x标志，因此它将退出并显示错误消息和非零返回代码
grep hello
*/
func main() {
	// 执行"date"命令，并捕获其输出
	//dateCmd执行了date命令，并打印了当前日期和时间
	//dateCmd := exec.Command("date")

	// 使用cmd.exe的/C选项来执行date命令
	dateCmd := exec.Command("cmd", "/C", "date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	//// 尝试执行"date -x"命令，但这次只处理错误
	//_, err = exec.Command("date", "-x").Output()
	//if err != nil {
	//	// 使用类型断言来检查错误的具体类型
	//	switch e := err.(type) {
	//	case *exec.Error:
	//		fmt.Println("failed executing:", err)
	//	case *exec.ExitError:
	//		// 如果命令退出时返回了非零状态码，则打印退出码
	//		fmt.Println("command exit rc =", e.ExitCode())
	//	default:
	//		// 对于其他类型的错误，使用panic来处理
	//		panic(err)
	//	}
	//}
	//// 执行"grep hello"命令，并向其标准输入写入数据
	////grepCmd执行了grep hello命令，并向其标准输入写入了两行文本。grep会过滤出包含"hello"的行
	//grepCmd := exec.Command("grep", "hello")
	//
	//grepIn, _ := grepCmd.StdinPipe()
	//grepOut, _ := grepCmd.StdoutPipe()
	//grepCmd.Start()
	//grepIn.Write([]byte("hello grep\ngoodbye grep"))
	//grepIn.Close()
	//grepBytes, _ := io.ReadAll(grepOut)
	//grepCmd.Wait()
	//
	//fmt.Println("> grep hello")
	//fmt.Println(string(grepBytes))
	//// 执行"bash -c 'ls -a -l -h'"命令，并捕获其输出
	//lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	//lsOut, err := lsCmd.Output()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("> ls -a -l -h")
	//fmt.Println(string(lsOut))
}
