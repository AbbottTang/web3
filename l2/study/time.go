package main

import (
	"fmt"
	"time"
)

/*
*

Go 对时间和持续时间提供了广泛的支持；这里有一些例子。
我们首先获取当前时间。
您可以time通过提供年、月、日等来构建结构。时间总是与Location时区相关联。
您可以按预期提取时间值的各个组成部分
周一至周日Weekday也可以。
这些方法比较两次，分别测试第一次是否发生在第二次之前、之后或同时发生。
该Sub方法返回Duration表示两个时间之间的间隔。
我们可以用各种单位来计算持续时间的长度。
您可以使用Add将时间提前指定的持续时间，或使用-将时间向后移动指定的持续时间。
*/
func main() {
	p := fmt.Println
	now := time.Now()
	p(now)
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

	// 假设我们有一个UTC时间
	utcTime := time.Date(2023, 4, 1, 12, 0, 0, 0, time.UTC)

	// 获取本地时区
	localLocation := time.Local

	// 将UTC时间转换为本地时间
	localTime := utcTime.In(localLocation)

	fmt.Println("UTC时间:", utcTime)
	fmt.Println("本地时间:", localTime)

	// 获取当前本地时间
	localTime2 := time.Now()

	// 将本地时间转换为UTC时间
	utcTime2 := localTime.UTC()

	fmt.Println("本地时间:", localTime2)
	fmt.Println("UTC时间:", utcTime2)
}
