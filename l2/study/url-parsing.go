package main

import (
	"fmt"
	"net"
	"net/url"
)

/*
*
url.Parse用于解析URL字符串，并返回一个*url.URL类型的值和一个错误。如果解析失败，错误将不为nil。

*url.URL类型有多个字段，如Scheme（协议方案）、User（用户信息）、Host（主机部分）、
Path（路径部分）、Fragment（片段部分）和RawQuery（原始查询字符串）等。

User字段是一个*url.Userinfo类型，它包含用户名和密码。可以使用Username和Password方法来分别获取用户名和密码。

net.SplitHostPort用于分割主机名和端口号。它返回一个主机名和一个端口号，以及一个错误（在正常情况下，错误为nil）。

url.ParseQuery用于解析查询字符串，并返回一个url.Values类型的map，其中键是查询参数的名字，
值是一个字符串切片，包含该参数的所有值。

在访问map的值时，需要注意检查键是否存在，以及切片是否为空。在上面的代码中，直接访问m["k"]是安全的，
因为已知查询字符串中包含k=v。但在实际应用中，应该添加适当的错误检查。
*/
func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析URL字符串
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 输出URL的协议方案
	fmt.Println(u.Scheme)
	// 输出: postgres

	// 输出URL的用户信息部分（包含用户名和可能存在的密码）
	fmt.Println(u.User)
	// 输出: user:pass

	// 输出URL中的用户名
	fmt.Println(u.User.Username())
	// 输出: user

	// 输出URL中的密码，忽略错误（在实际代码中应该检查错误）
	p, _ := u.User.Password()
	fmt.Println(p)
	// 输出: pass

	// 输出URL的主机部分（包含主机名和可能存在的端口号）
	fmt.Println(u.Host)
	// 输出: host.com:5432

	// 使用net.SplitHostPort分割主机名和端口号
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	// 输出: host.com

	fmt.Println(port)
	// 输出: 5432

	// 输出URL的路径部分
	fmt.Println(u.Path)
	// 输出: /path

	// 输出URL的片段（fragment）部分，即URL中的"#"后面的部分
	fmt.Println(u.Fragment)
	// 输出: f

	// 输出URL的原始查询字符串
	fmt.Println(u.RawQuery)
	// 输出: k=v

	// 解析查询字符串为map
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	// 输出: map[k:[v]]

	// 输出查询参数"k"的第一个值
	fmt.Println(m["k"])
	// 输出: v
}
