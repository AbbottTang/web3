package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
*
Go 提供对 JSON 编码和解码的内置支持，包括内置和自定义数据类型。
我们将使用这两个结构来演示下面自定义类型的编码和解码。
只有导出的字段才会以 JSON 格式进行编码/解码。字段必须以大写字母开头才能导出。
首先，我们将了解如何将基本数据类型编码为 JSON 字符串。以下是原子值的一些示例。
这里有一些切片和映射，它们按照您期望的方式编码为 JSON 数组和对象。
JSON 包可以自动编码您的自定义数据类型。它将只在编码输出中包含导出的字段，并默认使用这些名称作为 JSON 键
您可以在结构字段声明中使用标签来自定义编码的 JSON 键名。查看response2上面的定义以查看此类标签的示例。
现在让我们看看如何将 JSON 数据解码为 Go 值。这是一个通用数据结构的示例。
我们需要提供一个变量，JSON 包可以将解码的数据放入该变量中。
这 map[string]interface{}将保存字符串到任意数据类型的映射。
这是实际的解码和相关错误的检查。
为了使用解码映射中的值，我们需要将它们转换为适当的类型。例如，我们将其中的值转换num为预期的float64类型
访问嵌套数据需要一系列转换。
我们还可以将 JSON 解码为自定义数据类型。这样做的好处是可以为我们的程序增加额外的类型安全性，并且在访问解码的数据时无需进行类型断言。
在上面的例子中，我们总是使用字节和字符串作为数据和标准输出上的 JSON 表示之间的中间体。
我们还可以将 JSON 编码直接传输到os.Writers 类 os.Stdout或甚至 HTTP 响应主体。
*/

// 您定义了两个结构体response1和response2，其中response2使用了JSON标签来指定字段名。
type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

// json.Marshal 和 json.Unmarshal 在 Go 语言中是用于 JSON 序列化和反序列化的核心函数。
func main() {
	//您分别对布尔值、整数、浮点数、字符串、字符串切片以及映射（map）进行了JSON编码，并将编码后的字节切片转换为字符串以便打印。
	//这些操作展示了如何将Go语言中的基本数据类型转换为JSON格式的字符串。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page":1,"fruits":["apple","peach","pear"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//您创建了一个json.Encoder实例，用于将数据编码为JSON格式并输出到标准输出（os.Stdout）。
	//您使用Encoder的Encode方法将一个映射编码为JSON，并输出到控制台。
	//这展示了如何使用json.Encoder来编码数据并将其输出到io.Writer接口的实现中。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
