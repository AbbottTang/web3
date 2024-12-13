package main

import "fmt"

// 7*6*5*4*3*2*1*1
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	fmt.Println(fact(7))

	var fib func(n int) int
	//fib(7)=fib(6)+fib(5)=fib(5)+fib(4)+fib(4)+fib(3)
	//fib(3)=fib(2)+fib(1) =2
	//fib(1) = 1
	//fib(2) = 1
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
