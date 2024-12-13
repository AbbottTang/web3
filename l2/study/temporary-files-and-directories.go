package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check3(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//使用os.CreateTemp创建一个临时文件。该函数的第一个参数是目录前缀，如果为空字符串，
	//则系统会选择合适的临时文件目录。 第二个参数是文件名的前缀，
	//实际文件名会在这个前缀后面加上一些随机字符以避免冲突。函数返回一个文件句柄*os.File和一个错误值。
	f, err := os.CreateTemp("", "sample")
	check3(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check3(err)

	//使用os.MkdirTemp创建一个临时目录。该函数的参数与os.CreateTemp类似，
	//第一个参数是目录前缀，第二个参数是目录名的前缀。函数返回一个目录名称和一个错误值。
	dname, err := os.MkdirTemp("", "sampledir")
	check3(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)
	//使用filepath.Join将临时目录的名称和文件名file1连接起来，形成一个完整的文件路径。
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check3(err)
}
