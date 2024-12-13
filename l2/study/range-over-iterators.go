package main

import (
	"fmt"
	"iter"
	"slices"
)

//iter.Seq-yield讲解
///https://zhuanlan.zhihu.com/p/704474405
//同时编译多个go文件，多个狗文件中只能由一个main
// go run range-over-iterators.go .\generics.go
//type List[T any] struct {
//	head,tail *element[T]
//}

//type element[T any] struct {
//	next *element[T]
//	val  T
//}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	for e := range lst.All() {
		fmt.Println(e)
	}
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {
		fmt.Println(n)
		if n >= 900 {
			break
		}

	}
}
