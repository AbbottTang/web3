package main

/*
**
‌泛型类型参数‌：

S ~[]E：这表示 S 是一个切片类型，其元素类型为 E。~ 符号是Go语言中用于表示“类型约束”的，它指定了 S 必须是某种切片类型。
E comparable：这是一个类型约束，表示 E 必须是可比较的类型。可比较的类型包括所有基本类型（如 int、float64、string 等）以及实现了 comparable 接口的类型。
‌函数参数‌：

s S：这是传递给函数的切片，其类型为 S。
v E：这是要查找的元素，其类型为 E。
‌函数体‌：

for i := range s：这是一个循环，它遍历切片 s 中的所有元素。i 是当前元素的索引。
if v == s[i]：在每次循环中，检查切片中的当前元素 s[i] 是否等于要查找的元素 v。
return i：如果找到了匹配的元素，返回其索引 i。
return -1：如果循环结束后仍未找到匹配的元素，返回 -1。
这个函数非常有用，因为它可以处理任何类型的切片，只要其元素类型是可比较的。
*/
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// 在Go语言中，* 符号用于表示指针。指针是存储变量内存地址的变量，而不是直接存储值。通过指针，我们可以访问或修改存储在该地址上的值。
type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

/*
** & 符号用于获取变量的内存地址，也就是创建一个指向该变量的指针。
element[T] 是一个泛型结构体类型，它可能有一个名为 val 的字段（尽管在你之前的定义中，字段名是 value 而不是 val；我假设这里你是用 val 作为示例）。
{val: v} 是结构体字面量语法，用于初始化 element[T] 类型的一个新实例，并将 val 字段设置为 v。
& 符号前面是结构体字面量，因此它创建了一个新实例并返回该实例的内存地址（即指针）。
lst.tail.next 是一个指针字段，它指向链表中的下一个元素。通过将 lst.tail.next 设置为新创建的 element[T] 实例的地址，你将新元素添加到链表的末尾。
*/
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		//向后增加一个元素
		lst.tail.next = &element[T]{val: v}
		//tail指针向后移动
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

//func main() {
//	var s = []string{"foo", "var", "zoo"}
//	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
//	_ = SlicesIndex[[]string, string](s, "zoo")
//	var s1 = SlicesIndex[[]string, string](s, "zoo")
//	fmt.Println("index of zoo s1:", s1)
//	lst := List[int]{}
//	lst.Push(5)
//	lst.Push(7)
//	lst.Push(9)
//	lst.Push(10)
//	fmt.Println(lst.AllElements())
//}
