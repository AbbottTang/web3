// Go 内置了对正则表达式的支持。下面是 Go 中一些常见的与正则表达式相关的任务的示例。
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	//此段代码负责检查字符串"peach"是否与正则表达式p([a-z]+)ch相匹配，并返回布尔值结果。
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	//通过Compile函数，我们将正则表达式编译成一个Regexp对象，以便进行后续的多次匹配或查找操作，从而提高效率。
	r, _ := regexp.Compile("p([a-z]+)ch")
	//利用已编译的正则表达式对象r，我们可以便捷地检查字符串"peach"是否与正则表达式相匹配。

	fmt.Println(r.MatchString("peach"))
	//该方法用于在字符串"peach punch"中搜索第一个与正则表达式相匹配的子串。

	fmt.Println(r.FindString("peach punch"))
	//此功能返回匹配子串在原始字符串中的起始和结束索引位置。

	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	//该方法不仅返回匹配的整个子串，还包括所有括号内捕获的子组。

	fmt.Println(r.FindStringSubmatch("peach punch"))
	//与FindStringSubmatch相似，但返回的是各子匹配在原始字符串中的索引位置。

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	//在字符串"peach punch pinch"中查找所有与正则表达式匹配的子串，-1表示查找所有匹配项。

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	//‌查找所有子匹配字符串索引‌：返回所有匹配项及其子匹配的索引位置。
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	//‌限制查找数量‌：在字符串中仅查找前两个匹配的子串
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	//‌匹配字节切片‌：检查字节切片是否与正则表达式匹配。
	fmt.Println(r.Match([]byte("peach")))
	//‌使用MustCompile‌：MustCompile是Compile的封装，若正则表达式编译失败，则会导致程序崩溃。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)
	//‌替换字符串‌：将字符串中所有匹配的子串替换为"<fruit>"。
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	//自定义替换函数‌：利用ReplaceAllFunc方法，我们可以对匹配的子串执行自定义的替换逻辑，在此例中将匹配的子串转换为大写
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
