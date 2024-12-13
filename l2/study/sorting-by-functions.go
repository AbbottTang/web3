package main

import (
	"cmp"
	"fmt"
	"slices"
)

// 有时我们会想按集合的自然顺序以外的其他顺序对其进行排序。
// 例如，假设我们想按字符串的长度而不是按字母顺序对其进行排序。以下是 Go 中自定义排序的一个示例。/
// 我们实现了字符串长度的比较函数。cmp.Compare对此很有帮助。
// 现在我们可以调用slices.SortFunc这个自定义比较函数按名称长度排序fruits
// 我们可以使用相同的技术对非内置类型的值进行排序。
// people使用 按年龄排序slices.SortFunc。
// 注意：如果Person结构很大，您可能希望切片包含*Person并相应地调整排序函数。如有疑问，请进行基准测试！
func main() {
	fruits := []string{"apple", "orange", "banana", "kiwi"}
	lenCmp := func(i, j string) int {
		return cmp.Compare(len(i), len(j))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)
	type Person struct {
		name string
		age  int
	}
	people := []Person{
		{name: "jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)

}
