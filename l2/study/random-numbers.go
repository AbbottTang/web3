package main

import (
	"fmt"
	"math/rand/v2"
)

/**
rand.IntN(n) 生成一个 [0, n) 范围内的随机整数，即包括0但不包括n。

rand.Float64() 生成一个 [0.0, 1.0) 范围内的随机浮点数，即包括0.0但不包括1.0。

当您使用相同的种子（在这个例子中是42）初始化随机数生成器时，无论您创建多少次生成器，它们都会生成相同的随机数序列。这是因为随机数生成器是伪随机的，它们根据种子值来生成数列。

在实际的应用程序中，如果您需要不同的随机数序列，您应该使用不同的种子值来初始化随机数生成器。种子值通常来自于某个不可预测的来源，比如当前时间的纳秒数。

rand.NewPCG 和 rand.New 是用来创建自定义的随机数生成器和源的。在大多数情况下，使用默认的随机数生成器（通过直接调用 rand 包中的函数）就足够了。
*/

func main() {
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
