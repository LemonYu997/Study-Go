package main

import "fmt"

// 数组的长度不同可以认为是不同类型
// 数组是值传递
func main() {
	// 这两个是不同的类型！
	var a [2][3]int
	var b [3][2]int

	// 以下代码会编译错误
	// a = b  // 错误：类型不匹配

	fmt.Printf("a 的类型: %T\n", a) // [2][3]int
	fmt.Printf("b 的类型: %T\n", b) // [3][2]int

	fmt.Println("---------------------------------")

	// 数组是值传递

	// 数组是值类型
	original := [2][2]int{{1, 2}, {3, 4}}

	// 赋值会创建副本
	copy := original

	// 修改副本不会影响原始数组
	copy[0][0] = 100

	fmt.Println("原始数组:", original) // [[1 2] [3 4]]
	fmt.Println("副本数组:", copy)     // [[100 2] [3 4]]

	// 如果需要引用语义，可以使用切片（后续文章会介绍）
}
