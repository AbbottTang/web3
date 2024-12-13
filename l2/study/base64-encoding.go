package main

import (
	b64 "encoding/base64"
	"fmt"
)

/*
*
b64.StdEncoding是标准Base64编码，它使用+、/和=作为特殊字符。这种编码在大多数场景下都是有效的，但在URL或文件名中可能会遇到问题，因为这些特殊字符可能需要被编码或转义。

b64.URLEncoding是URL兼容的Base64编码，它使用-、_来替换+、/，并且省略了尾部的=填充字符（如果可能的话）。这种编码更适合在URL或文件名中使用。

在解码时，无论使用哪种编码方式，DecodeString函数都会返回一个字节切片和一个错误。在实际代码中，您应该检查这个错误是否为nil，以确保解码成功。在上述示例中，错误被忽略了（使用_空白标识符），这在实际应用中通常是不推荐的。

Base64编码是一种将二进制数据转换为ASCII字符串的方式，它通常用于在文本协议（如HTTP、MIME等）中传输二进制数据。编码后的数据会比原始数据稍大（大约增加33%），但这是为了能够在文本环境中安全地传输二进制数据所必需的。
*/
func main() {
	data := "abc123!?$*&()'-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
