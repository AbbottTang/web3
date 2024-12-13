package main

import (
	"fmt"
	"time"
)

/*
*

Go 支持通过基于模式的布局进行时间格式化和解析。
这是根据 RFC3339 使用相应的布局常量格式化时间的基本示例。
时间解析使用与相同的布局值Format。
Format并使用基于示例的布局。通常，您会为这些布局Parse使用常量，但您也可以提供自定义布局。
布局必须使用参考时间来显示格式化/解析给定时间/字符串的模式。
示例时间必须与所示完全相同：2006 年、15 为小时、星期一为星期几等。timeMon Jan 2 15:04:05 MST 2006
对于纯数字表示，您还可以将标准字符串格式与时间值的提取部分一起使用。
Parse将返回有关格式错误的输入的错误，解释解析问题。
*/
func main() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(time.RFC3339, "2018-08-08T00:00:00Z")

	p(t1)
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05Z07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
