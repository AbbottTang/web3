package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}
func plusPlus(a, b, c int) int {
	return a + b + c
}
func main() {
	res := plus(1, 2)
	fmt.Println(res)
	res = plusPlus(4, 5, 6)
	fmt.Println(res)
}
