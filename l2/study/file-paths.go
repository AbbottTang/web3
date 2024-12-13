package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	//filepath.Join：用于将多个路径片段连接成一个路径。它会自动处理路径中的冗余分隔符和相对路径。
	//p := filepath.Join("dir1", "dir2", "filename") 将生成路径 dir1/dir2/filename（在Unix-like系统上）或 dir1\dir2\filename（在Windows上）。
	//filepath.Join("dir1//", "filename") 会自动去除多余的斜杠，生成 dir1/filename。
	//filepath.Join("dir1/../dir1", "filename") 会处理相对路径，.. 表示上一级目录，因此这里相当于 dir1/filename。
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	//filepath.Dir：返回路径中的目录部分。
	//filepath.Dir(p) 将返回 dir1/dir2（或相应的Windows路径）。
	//filepath.Base：返回路径中的文件名部分。
	//filepath.Base(p) 将返回 filename。
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	//filepath.IsAbs：判断路径是否为绝对路径。
	//filepath.IsAbs("dir/file") 将返回 false，因为它是一个相对路径。
	//filepath.IsAbs("/dir/file") 将返回 true（在Unix-like系统上），因为它是一个绝对路径。在Windows上，绝对路径通常以盘符（如 C:\）开头。
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	//filepath.Ext：返回文件路径中的扩展名部分。
	//filepath.Ext(filename) 对于 filename := "config.json" 将返回 .json。
	//strings.TrimSuffix：从字符串末尾去除指定的后缀。
	//strings.TrimSuffix(filename, ext) 将返回没有扩展名的文件名，即 config。
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))
	//filepath.Rel：返回将basepath相对路径转换为与targpath相对的路径。如果basepath和targpath有共同的前缀，这个函数会返回targpath中剩余的部分。
	//第一个 filepath.Rel("a/b", "a/b/t/file") 调用将返回 t/file，因为a/b是a/b/t/file的共同前缀。
	//第二个 filepath.Rel("a/b", "a/c/t/file") 调用将返回 ../c/t/file，因为a/b和a/c/t/file的共同前缀是a/，需要从a/c/t/file回退到上一级目录（..）再进入c/t/file
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
