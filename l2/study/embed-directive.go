package main

//这段Go语言代码使用了embed包来在编译时将文件或文件夹嵌入到程序中。
//embed包允许你在Go程序中直接包含文件内容，这样在运行时就不需要依赖外部文件系统来访问这些文件了。

//go:embed是一个编译器指令，允许程序在构建时将任意文件和文件夹包含到 Go 二进制文件中。在

import (
	"embed"
)

// //go:embed folder/single_file.txt 指令将folder/single_file.txt文件的内容嵌入到程序中，
// 并绑定到fileString变量。
// 由于fileString被声明为string类型，因此文件内容会被解码为字符串。
//
//	2、嵌入文件只能为源码文件同级目录和子目录下的文件。
//
//go:embed folder/single_file.txt
var fileString string

// //go:embed folder/single_file.txt 指令再次被使用，但这次是将文件内容绑定到fileByte变量，
// 由于fileByte被声明为[]byte类型，因此文件内容会被解码为字节切片。
//
//go:embed folder/single_file.txt
var fileByte []byte

// //go:embed folder/*.hash 指令将folder目录下所有以.hash为扩展名的文件嵌入到程序中，
// 并绑定到folder变量，folder的类型是embed.FS，它是一个表示嵌入文件系统的类型。
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
