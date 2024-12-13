package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check2(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//使用os.Mkdir创建一个名为subdir的目录，权限设置为0755（可读可写可执行由所有者，可读可执行由组和其他用户）。
	err := os.Mkdir("C:/Users/Administrator/Desktop/tmp/subdir", 0755)
	check2(err)
	//使用defer os.RemoveAll("subdir")确保在main函数结束时删除subdir及其所有内容，无论中间是否发生错误。
	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check2(os.WriteFile(name, d, 0644))
	}
	//使用createEmptyFile在subdir及其子目录中创建几个空文件。
	createEmptyFile("C:/Users/Administrator/Desktop/tmp/subdir/file1")
	//使用os.MkdirAll创建嵌套的目录subdir/parent/child，权限同样设置为0755。这个函数会创建所有必要的父目录。
	err = os.MkdirAll("C:/Users/Administrator/Desktop/tmp/subdir/parent/child", 0755)
	check2(err)

	createEmptyFile("C:/Users/Administrator/Desktop/tmp/subdir/parent/file2")
	createEmptyFile("C:/Users/Administrator/Desktop/tmp/subdir/parent/file3")
	createEmptyFile("C:/Users/Administrator/Desktop/tmp/subdir/parent/child/file4")
	//使用os.ReadDir列出subdir/parent目录的内容，并打印每个条目的名称和是否为目录。
	c, err := os.ReadDir("C:/Users/Administrator/Desktop/tmp/subdir/parent")
	check2(err)

	fmt.Println("Listing C:/Users/Administrator/Desktop/tmp/subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	//使用os.Chdir改变当前工作目录到subdir/parent/child，然后再次使用os.ReadDir列出当前目录（.）的内容。
	err = os.Chdir("C:/Users/Administrator/Desktop/tmp/subdir/parent/child")
	check2(err)

	c, err = os.ReadDir(".")
	check2(err)

	fmt.Println("Listing C:/Users/Administrator/Desktop/tmp/subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	//再次使用os.Chdir改变当前工作目录回到subdir的父目录（通过相对路径"../../.."）。
	err = os.Chdir("../../..")
	check2(err)
	//使用filepath.WalkDir遍历subdir目录树，对每个条目调用visit函数。
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

// visit是一个回调函数，用于filepath.WalkDir。它接收路径、目录条目和错误作为参数。
// 如果错误不为nil，则直接返回错误。
// 否则，打印路径和该条目是否为目录。
// 返回nil表示没有错误，继续遍历。
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
