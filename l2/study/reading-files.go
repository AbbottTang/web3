package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
在您的Go代码示例中，您展示了如何使用os和bufio包来读取文件内容、移动文件指针（seek）以及使用缓冲读取器（bufio.Reader）。
以下是代码的解释以及潜在的输出（假设文件/tmp/dat存在并且包含足够的数据）：
*/
// 错误检查函数，如果发生错误则触发panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 使用os.ReadFile读取整个文件内容
	dat, err := os.ReadFile("C:\\Users\\Administrator\\Desktop\\tmp/dat.txt")
	check(err)
	fmt.Print(string(dat)) // 输出文件内容

	// 使用os.Open打开文件，准备进行更细粒度的读取
	f, err := os.Open("C:\\Users\\Administrator\\Desktop\\tmp/dat.txt")
	check(err)
	defer f.Close() // 确保在main函数结束时关闭文件

	// 读取文件的前5个字节
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1])) // 输出读取的字节数和内容

	// 将文件指针移动到文件的第6个字节（从文件开始计算）
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	// 读取接下来的2个字节
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2[:n2])) // 输出读取的字节数、文件指针位置和读取的内容

	// 将文件指针向前移动4个字节（相对于当前位置）
	_, err = f.Seek(4, io.SeekCurrent)
	check(err)
	// 将文件指针向后移动10个字节（相对于文件末尾）
	_, err = f.Seek(-10, io.SeekEnd)
	check(err)
	// 再次将文件指针移动到文件的第6个字节（从文件开始计算）
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	// 使用io.ReadAtLeast读取至少2个字节
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3)) // 输出读取的字节数、文件指针位置和读取的内容

	// 将文件指针移动到文件开始位置
	_, err = f.Seek(0, io.SeekStart)
	check(err)
	// 创建一个缓冲读取器
	r4 := bufio.NewReader(f)
	// 使用Peek方法查看但不消耗接下来的5个字节
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4)) // 输出查看的字节内容

	// 关闭文件（由于使用了defer，这里实际上不需要再次调用Close）
	f.Close()
}

/**
os.ReadFile会读取整个文件的内容，并将其作为一个字节切片返回。这对于小文件来说很方便，但对于大文件可能会导致内存使用过高。

os.Open打开一个文件，返回一个*os.File对象，该对象可以用于更细粒度的读取和写入操作。

f.Read方法从文件中读取数据到一个字节切片中，并返回读取的字节数和一个错误。如果切片长度小于文件剩余内容，则只读取切片长度的数据。

f.Seek方法用于移动文件指针到指定的位置。io.SeekStart、io.SeekCurrent和io.SeekEnd分别表示从文件开始、当前位置和文件末尾开始计算偏移。

io.ReadAtLeast函数读取至少指定数量的字节，如果读取的字节数少于指定数量，则会返回一个错误。

bufio.NewReader创建一个缓冲读取器，它可以提高读取效率，并提供一些额外的读取方法，如Peek，它允许你查看但不消耗输入。

使用defer语句确保文件在main函数结束时被关闭，这是一种良好的实践，可以防止文件句柄泄漏。

在实际应用中，应该始终检查错误并适当处理，而不是简单地使用check函数触发panic。这里使用check函数只是为了简化示例代码。
*/
