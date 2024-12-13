//字符串函数

//标准库的strings包提供了很多有用的字符串相关函数。下面是一些示例，让您了解该包
//我们将其别名fmt.Println为一个较短的名称，因为我们将在下面多次使用它。
//以下是strings 中可用的函数示例 。由于这些是来自包的函数，而不是字符串对象本身的方法，
//因此我们需要将相关字符串作为第一个参数传递给函数。您可以在strings 包文档中找到更多函数

package main

import (
	"fmt"
	str "strings"
)

var p = fmt.Println

func main() {

	p("Contains:  ", str.Contains("test", "es"))
	p("Count:     ", str.Count("test", "t"))
	p("HasPrefix: ", str.HasPrefix("test", "te"))
	p("HasSuffix: ", str.HasSuffix("test", "st"))
	p("Index:     ", str.Index("test", "e"))
	p("Join:      ", str.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", str.Repeat("a", 5))
	p("Replace:   ", str.Replace("foo", "o", "0", -1))
	p("Replace:   ", str.Replace("foo", "o", "0", 1))
	p("Split:     ", str.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", str.ToLower("TEST"))
	p("ToUpper:   ", str.ToUpper("test"))
}
