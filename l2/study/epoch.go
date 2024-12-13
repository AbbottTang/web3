package main

import (
	"fmt"
	"time"
)

/*
*

程序中常见的需求是获取自 Unix 纪元以来的秒数、毫秒数或纳秒数。以下是如何在 Go 中实现此目的。
使用、time.Now或 分别获取自 Unix 纪元以来经过的时间（以秒、毫秒或纳秒为单位）。Unix,UnixMilli,UnixNano
您还可以将自纪元以来的整数秒数或纳秒数转换为相应的time。
接下来我们来看看另一个与时间相关的任务：时间解析和格式化。
*/
func main() {
	// 获取当前本地时间
	now := time.Now()
	// 输出当前本地时间的字符串表示，例如：2023-04-01 12:34:56.789123456 +0800 CST m=+0.000000000
	fmt.Println(now)

	// 获取当前时间的Unix时间戳（秒）
	// Unix时间戳是自1970年1月1日00:00:00 UTC以来的秒数
	// 输出一个整数，例如：1677650096
	fmt.Println(now.Unix())
	// 获取当前时间的Unix时间戳（毫秒）
	// 注意：UnixMilli并不是一个标准的方法，这可能是您自定义的或某个特定版本的Go中的扩展方法。
	// 标准Go库中没有直接提供UnixMilli方法，但您可以通过now.UnixNano() / 1e6来计算毫秒数。
	// 假设这里它是一个有效的方法，它会输出一个整数，例如：1677650096789
	// 这行代码可能不会编译，除非您自己定义了UnixMilli方法
	fmt.Println(now.UnixMilli())
	// 获取当前时间的Unix时间戳（纳秒）
	// UnixNano返回的是自1970年1月1日00:00:00 UTC以来的纳秒数
	// 输出一个整数，例如：1677650096789123456
	fmt.Println(now.UnixNano())
	// 使用Unix方法根据秒和纳秒创建时间
	// 这里使用now.Unix()作为秒数，0作为纳秒数
	// 这将返回一个新的time.Time对象，表示与now相同的秒数，但纳秒部分被设置为0
	fmt.Println(time.Unix(now.Unix(), 0))
	// 使用Unix方法根据纳秒创建时间
	// 这里使用now.UnixNano()作为纳秒数，0作为秒数（实际上这是不常见的用法，因为通常秒数不会为0）
	// 这将返回一个新的time.Time对象，表示从1970年1月1日00:00:00 UTC开始的now.UnixNano()纳秒后的时间
	fmt.Println(time.Unix(0, now.UnixMilli()))

}
