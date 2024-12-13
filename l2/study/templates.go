package main

import (
	"html/template"
	"os"
)

//Go 提供内置支持，可使用此text/template包创建动态内容或向用户显示自定义输出。
//名为 的兄弟包html/template提供相同的 API，但具有额外的安全功能，应该用于生成 HTML。
//我们可以创建一个新模板并从字符串中解析其主体。模板是静态文本和其中包含的 {{...}}用于动态插入内容的“操作”的混合体
//或者，我们可以使用函数来在返回错误template.Must时 panic 。这对于在全局范围内初始化的模板特别有用。Parse
//通过“执行”模板，我们生成其文本，其中包含其操作的具体值。{{.}}操作由作为参数传递给 的值替换Execute。
//下面我们将使用辅助函数
//如果数据是结构体，我们可以使用{{.FieldName}}操作来访问其字段。模板执行时，字段应导出以便访问。
//这同样适用于map；,map对键名称的大小写没有限制。
//if/else 为模板提供条件执行。如果某个值是某个类型的默认值（例如 0、空字符串、nil 指针等），则该值被视为 false。
//此示例演示了模板的另一个功能：-在操作中使用来修剪空格。
//范围块让我们可以循环遍历切片、数组、映射或通道。范围块内部{{.}}设置为迭代的当前项。

func main() {
	//‌创建和解析模板‌，您首先创建了一个名为t1的新模板，
	//并解析了一个简单的模板字符串，该字符串包含一个占位符{{.}}，用于插入动态内容。
	t1 := template.New("t1")
	t1, err := t1.Parse("value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	//‌使用Must函数‌，Must函数用于简化模板解析过程中的错误处理。如果解析失败，Must会触发panic
	t1 = template.Must(t1.Parse("value:{{.}}\n"))
	//您使用Execute方法将模板与数据结合，并将结果输出到标准输出（os.Stdout）。
	//这里，您分别传入了字符串和整数作为数据。
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})
	//创建模板的辅助函数‌：定义了一个名为Create的辅助函数，用于简化模板的创建和解析过程。
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	//使用Create函数创建了一个新模板t2，并使用结构体和映射作为数据来执行模板。
	t2 := Create("t2", "Name:{{.name}}\n")
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})
	//创建了一个包含条件判断的模板t3。如果传入的数据非空，则输出"yes"；否则，输出"no"。
	// -在操作中使用来修剪空格
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")
	//创建了一个包含范围循环的模板t4，用于遍历并输出一个字符串切片中的每个元素。
	t4 := Create("t4", "Range:{{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})

}
