package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// 在Go语言中，* 符号用于表示指针。指针是存储变量内存地址的变量，而不是直接存储值。通过指针，我们可以访问或修改存储在该地址上的值。
// & 符号用于获取变量的内存地址，也就是创建一个指向该变量的指针。
// errors.As是 的更高级版本errors.Is。它检查给定的错误（或其链中的任何错误）是否与特定错误类型匹配，并转换为该类型的值，返回true。如果不匹配，则返回false
func main() {

	_, err := f2(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
