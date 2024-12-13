package main

import (
	"encoding/xml"
	"fmt"
)

/*
*

Go 通过包提供对 XML 和类似 XML 格式的内置支持encoding/xml。
Plant 将映射到 XML。与 JSON 示例类似，字段标签包含编码器和解码器的指令。
这里我们使用 XML 包的一些特殊功能：字段XMLName名称指示表示此结构的 XML 元素的名称；
id,attr表示该Id字段是 XML 属性而不是嵌套元素。
发出代表我们工厂的 XML；用于 MarshalIndent产生更易于人类阅读的输出。
要将通用 XML 标头添加到输出，请明确附加它。
用于Unmarshal将 XML 字节流解析为数据结构。如果 XML 格式错误或无法映射到 Plant，则会返回描述性错误。
字段标签parent>child>plant告诉编码器将所有的plants 嵌套在<parent><child>...
*/
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v name=%v origin=%v", p.Id, p.Name, p.Origin)

}
func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
