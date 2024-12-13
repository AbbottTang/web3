package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("key:", k, " value:", v)
	}
	for K := range kvs {
		fmt.Println("key:", K)
	}
	for i, c := range "hello world" {
		fmt.Println(i, c)
	}

}
