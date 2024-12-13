package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	fmt.Printf("struct2: %+v\n", p)

	fmt.Printf("struct3: %#v\n", p)

	fmt.Printf("type: %T\n", p)

	fmt.Printf("bool: %t\n", true)

	fmt.Printf("int: %d\n", 123)

	fmt.Printf("bin: %b\n", 14)

	fmt.Printf("char: %c\n", 33)

	fmt.Printf("hex: %x\n", 456)

	fmt.Printf("float1: %f\n", 78.9)

	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"")

	fmt.Printf("str2: %q\n", "\"string\"")

	fmt.Printf("str3: %x\n", "hex this")

	fmt.Printf("pointer: %p\n", &p)

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}

/**
%v：以默认格式打印结构体。
%+v：在打印结构体时，会添加字段名。
%#v：打印结构体的Go语法表示（更详细）。
%T用于打印变量的类型。

‌打印基本数据类型‌：

布尔值、整数、二进制、字符、十六进制、浮点数等，分别使用%t、%d、%b、%c、%x（或%X用于大写）、%f、%e（或%E用于大写）等格式化动词。
‌打印字符串‌：

%s：打印字符串。
%q：打印带引号的字符串。
对于字符串的十六进制表示，可以使用%x（但通常需要对字符串进行转换）。
打印指针‌：%p用于打印指针的值。
‌宽度与精度控制‌：

通过在格式化动词前添加数字，可以控制输出的宽度和精度（例如，%6d表示宽度为6的整数，%6.2f表示宽度为6、精度为2的浮点数）。
使用-标志可以左对齐输出。
‌使用fmt.Sprintf和fmt.Fprintf‌：

fmt.Sprintf：将格式化的字符串返回，而不是打印到标准输出。
fmt.Fprintf：将格式化的字符串打印到指定的io.Writer（在此例中为os.Stderr）
*/
