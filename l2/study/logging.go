/*
*

Go 标准库提供了用于从 Go 程序输出日志的简单工具，其中log包用于自由格式输出， log/slog包用于结构化输出
Println只需从 包中调用诸如这样的函数即可log使用标准记录器，该记录器已预先配置为合理的日志输出到。
诸如或之os.Stderr类的附加方法 将在记录后退出程序。Fatal*Panic*

可以使用标志配置记录器以设置其输出格式。
默认情况下，标准记录器设置了log.Ldate和log.Ltime标志，这些标志收集在中log.LstdFlags。
例如，我们可以更改其标志以微秒精度发出时间。

它还支持发出log调用该函数的文件名和行。

创建自定义记录器并传递它可能会很有用。创建新记录器时，我们可以设置前缀以将其输出与其他记录器区分开来。

我们可以使用该方法在现有的记录器（包括标准记录器）上设置前缀SetPrefix。
记录器可以有自定义的输出目标；任何io.Writer作品
此调用将日志输出写入buf
这实际上会将其显示在标准输出上。

该slog包提供 结构化的日志输出。例如，以 JSON 格式记录日志非常简单
除了消息之外，slog输出还可以包含任意数量的键=值对。
示例输出；发出的日期和时间取决于示例的运行时间。
为了在网站上清晰呈现，这些内容被包装起来；实际上，它们是在一行上发出的。
*/
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {
	// 使用默认的logger打印一条消息
	log.Println("standard logger")

	// 设置logger的标志以包含标准标志和微秒
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// 设置logger的标志以包含标准标志和短文件路径/行号
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// 创建一个新的logger，使用os.Stdout作为输出，前缀为"my:"，并包含标准标志
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// 更改新logger的前缀为"ohmy:"
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// 创建一个缓冲区，并使用它创建一个新的logger
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")
	// 打印缓冲区中的内容，显示来自buflog的消息
	fmt.Print("from buflog:", buf.String())

	// 创建一个JSON格式的日志处理器，并将错误输出到os.Stderr
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	// 使用JSON处理器创建一个新的slog logger
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	// 打印带有额外键值对的信息
	myslog.Info("hello again", "key", "val", "age", 25)
	myslog.Error("key", "val")
}
