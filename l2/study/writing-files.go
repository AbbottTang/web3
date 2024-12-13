package main

import (
	"bufio"
	"fmt"
	"os"
)

// 错误检查函数，如果发生错误则触发panic
func checkWrite(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 准备要写入文件的数据
	d1 := []byte("hello\ngo\n")

	// 使用os.WriteFile将数据写入/tmp/dat1文件
	// 第三个参数是文件权限，0644表示所有者可以读写，组和其他用户可以读
	err := os.WriteFile("C:\\Users\\Administrator\\Desktop/tmp/dat1", d1, 0644)
	checkWrite(err)

	// 使用os.Create创建/tmp/dat2文件，如果文件已存在则会被截断
	f, err := os.Create("C:\\Users\\Administrator\\Desktop\\tmp/dat2")
	checkWrite(err)
	defer f.Close() // 确保在main函数结束时关闭文件

	// 准备要写入文件的字节数据
	d2 := []byte{115, 111, 109, 101, 10} // 对应字符串 "some\n"
	// 将数据写入文件
	n2, err := f.Write(d2)
	checkWrite(err)
	fmt.Printf("wrote %d bytes\n", n2) // 输出: wrote 5 bytes

	// 使用WriteString方法将字符串写入文件
	n3, err := f.WriteString("writes\n")
	checkWrite(err)
	fmt.Printf("wrote %d bytes\n", n3) // 输出: wrote 7 bytes

	// Sync方法将缓冲区中的数据写入底层文件，确保数据持久化
	f.Sync()

	// 创建一个缓冲写入器
	w := bufio.NewWriter(f)
	// 使用缓冲写入器的WriteString方法写入数据
	n4, err := w.WriteString("buffered\n")
	checkWrite(err)
	fmt.Printf("wrote %d bytes\n", n4) // 输出: wrote 9 bytes

	// Flush方法将缓冲写入器中的数据写入底层文件
	w.Flush()
}

/**
os.WriteFile会创建或截断文件，并将数据写入其中。如果文件已存在，其内容会被新数据覆盖。

os.Create会创建文件，如果文件已存在，则会截断（清空）文件。它返回一个*os.File对象，该对象可以用于后续的写入操作。

f.Write和f.WriteString方法直接将数据写入文件。Write方法接受一个字节切片作为参数，
而WriteString方法接受一个字符串作为参数。

f.Sync方法将文件描述符的缓冲区中的数据刷新到硬盘上，确保数据的持久化。
这对于确保数据在程序崩溃或系统崩溃后不会丢失很重要。

bufio.NewWriter创建一个缓冲写入器，它可以提高写入效率。
缓冲写入器会先将数据写入内存缓冲区，当缓冲区满或调用Flush方法时，再将数据写入文件。

w.WriteString方法将字符串写入缓冲写入器。由于使用了缓冲，所以这个方法可能会比直接写入文件更快，
但数据不会立即写入文件，而是先写入缓冲区。

w.Flush方法将缓冲写入器中的数据写入文件。在调用Flush之前，数据可能只存在于内存缓冲区中。

使用defer语句确保文件在main函数结束时被关闭，这是一种良好的实践，可以防止文件句柄泄漏和未写入的数据丢失。

在实际应用中，应该始终检查错误并适当处理，而不是简单地使用check函数触发panic。
这里使用check函数只是为了简化示例代码。
*/
