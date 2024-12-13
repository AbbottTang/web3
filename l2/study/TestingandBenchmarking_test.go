package main

import (
	"fmt"
	"testing"
)

// 代码示例包含了单元测试、基准测试以及一个被测试的函数IntMin。
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 基本测试 这个测试检查IntMin(2, -2)是否返回-2，如果不是，则通过t.Errorf记录一个错误。
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {

		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// 表驱动测试使用了一个结构体切片来定义多个测试用例，每个测试用例都包含输入参数a和b以及期望的结果want。
// 然后，它遍历这些测试用例，并对每个测试用例运行一个子测试。
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// 基准测试用于测量代码的性能。这个基准测试测量调用IntMin(1, 2)的性能。
// b.N是由基准测试框架提供的，表示应该运行测试循环的次数，以确保结果的准确性。
func BenchmarkIntMin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

//go test .\TestingandBenchmarking_test.go -v -run=TestIntMinBasic
//go test .\TestingandBenchmarking_test.go -v -run=TestIntMinTableDriven
//go test .\TestingandBenchmarking_test.go  -bench=.  -run=BenchmarkIntMin

//‌导入包‌：你的代码中同时导入了fmt和testing包，这是正确的，因为你在测试中使用了fmt来格式化字符串，而testing包提供了测试和基准测试的功能。
//
//‌测试运行‌：要运行这些测试，你可以使用go test命令。要运行基准测试，你可以使用go test -bench=.命令。
//
//‌代码组织‌：通常，测试代码会被放在与被测试代码相同的包中，但在不同的文件中，例如*_test.go。
//
//‌错误处理‌：在单元测试中，你使用了t.Errorf来记录错误。这是正确的做法，因为当测试失败时，它会提供有用的反馈。
//
//‌性能‌：基准测试应该关注你关心的性能点，并且应该使用代表性的输入来模拟真实世界的使用情况

//在Go语言中，若你只想运行特定的测试，可以使用go test命令结合-run标志来实现。-run标志后面需跟一个正则表达式，Go测试工具会依据此表达式来匹配并决定执行哪些测试函数。
//
//举例来说，假设你有一系列测试函数，如TestExample1、TestExample2等，而你只想运行TestExample1，那么可以在命令行中输入以下指令：
//go test -v -run=TestExample1

//当你在命令行中运行 go test 但收到错误信息 ? command-line-arguments [no test files] 时，这通常意味着 Go 测试工具在当前目录或其子目录中未能找到任何测试文件。测试文件是指那些包含 Test、Benchmark 或 Example 函数，并且文件名以 _test.go 结尾的 Go 源文件。
//
//出现这个错误的原因可能有以下几种：
//
//‌当前目录中没有测试文件‌：你可能不在一个包含测试文件的目录中，或者测试文件的命名不符合 Go 的约定（即文件名未以 _test.go 结尾）。
//
//‌包路径问题‌：如果你在一个子目录中运行 go test，但该子目录没有自己的 package 声明，或者 package 声明与目录结构不匹配，Go 可能会找不到测试文件。确保你的测试文件位于正确的包路径下，并且 package 声明与目录结构相对应。
//
//‌测试函数命名问题‌：测试函数的名称必须以 Test 开头，并且是可导出的（即首字母大写）。如果测试函数的名称不符合这个约定，Go 将不会识别它们为测试函数。
//
//‌使用了错误的目录或文件‌：你可能在错误的目录中运行了 go test，或者尝试测试的文件不包含任何测试函数。
//
//‌环境变量或 Go 配置问题‌：在某些情况下，环境变量（如 GOPATH）或 Go 的配置可能会影响测试文件的查找。确保你的 Go 环境配置正确。
//
//为了解决这个问题，你可以：
//
//检查当前目录及其子目录中是否存在以 _test.go 结尾的文件。
//确保测试文件的 package 声明与目录结构相匹配。
//检查测试函数的名称是否符合 Go 的约定（即以 Test 开头，首字母大写）。
//在正确的目录中运行 go test 命令。
//如果你的测试文件位于特定的包中，请确保你使用了正确的包路径来运行 go test。例如，如果测试文件位于 mypackage 包中，你应该在包含该包的目录中运行 go test mypackage。
//如果以上步骤都无法解决问题，你可能需要更详细地检查你的项目结构和 Go 环境配置。
