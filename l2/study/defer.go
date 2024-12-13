package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("C:/Users/Administrator/Desktop/tmp/test.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data hello world")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error closing file:%v\n", err)
		os.Exit(1)
	}
}
