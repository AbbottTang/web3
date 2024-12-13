// 数字解析
package main

import (
	"fmt"
	"strconv"
)

/*
*
在使用strconv.ParseFloat、strconv.ParseInt和strconv.ParseUint时，第二个参数表示要解析的数的位数（32或64）。
对于ParseFloat，它还可以是strconv.ParseFloat32来解析为float32类型。

strconv.ParseInt和strconv.ParseUint的第三个参数是基数（radix），它可以是2到36之间的数，或者是0。
当基数为0时，函数会尝试自动检测基数，这允许你解析十进制、十六进制（以0x或0X开头）等格式的数。

strconv.Atoi是strconv.ParseInt的一个简便封装，它只解析十进制的有符号整数，并且返回的是int类型而不是int64。

在解析字符串为数值时，如果字符串不是有效的数值表示，strconv的函数会返回一个错误。
在实际代码中，你应该总是检查这个错误是否为nil，以确保解析成功。
在你的示例中，错误被忽略了（使用_空白标识符），这在实际应用中通常是不推荐的。

当解析十六进制数时，确保字符串以0x或0X开头，否则strconv.ParseInt可能无法正确解析它。
*/
func main() {
	// ParseFloat将字符串解析为浮点数
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	// 输出: 1.234

	// ParseInt将字符串解析为有符号整数，基数为0表示自动检测（十进制、十六进制等）
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	// 输出: 123

	// ParseInt也可以解析十六进制数，这里"0x1c8"是十六进制的456
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	// 输出: 456

	// ParseUint将字符串解析为无符号整数，基数同样为0表示自动检测
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	// 输出: 789

	// Atoi是ParseInt的简便函数，用于解析十进制的有符号整数
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	// 输出: 135

	// 尝试使用Atoi解析非数字字符串，将返回错误
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
	// 输出: strconv.Atoi: parsing "wat": invalid syntax
}
