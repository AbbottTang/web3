package main

import (
	"fmt"
	"os"
)

/*
*
defer 关键字‌：defer 语句会延迟函数的执行，直到包含它的函数（在这个例子中是 main）返回为止。
但是，如果函数在 defer 语句之后调用了 os.Exit、panic 或发生了其他导致程序立即终止的情况，
那么 defer 语句中的函数就不会被执行。

‌os.Exit 函数‌：os.Exit 函数会立即终止程序的执行，并返回一个整数值作为程序的退出状态码。
这个状态码可以被操作系统或其他程序用来判断程序是否成功执行。在您的代码中，os.Exit(3) 表示程序以状态码 3 退出，
这通常表示程序遇到了某种错误或异常情况。

执行结果
当您运行这段代码时，程序会立即终止，并且不会打印出 !。
这是因为 os.Exit(3) 语句在 defer fmt.Println("!") 语句之后被执行，
导致程序在 defer 语句有机会执行之前就已经终止了。

注意事项
在使用 defer 时，要注意它的执行顺序。defer 语句会按照它们出现的顺序的相反顺序执行。
但是，如果程序在 defer 语句有机会执行之前就已经终止了（比如因为调用了 os.Exit），
那么这些 defer 语句就不会被执行。
os.Exit 函数会立即终止程序，并且不会执行任何在之后安排的 defer 语句或其他的清理代码。
因此，在使用 os.Exit 时要特别小心，确保在调用它之前已经完成了所有必要的清理工作。
*/
func main() {
	// defer 语句会安排紧随其后的函数调用（在这个例子中是 fmt.Println("!")）
	// 在 main 函数返回之前执行。但是，由于 os.Exit 会立即终止程序，
	// 所以这里的 defer 语句实际上不会被执行。
	defer fmt.Println("!")
	// os.Exit(3) 会立即终止程序，并向操作系统返回状态码 3。
	// 这个状态码可以被操作系统或调用此程序的其他程序用来判断程序是否正常结束。
	// 通常，状态码 0 表示成功，非 0 表示出现了某种错误或异常情况。
	os.Exit(3)
}

//$?只会保存上一个命令的退出状态码。如果你运行了多个命令而不检查中间的状态码，那么$?只会反映最后一个命令的结果。
