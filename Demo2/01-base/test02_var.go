package main

import "fmt"

// 全局变量
//var (
//	h int
//	i bool = false
//)

func main() {
	var a string = "hello"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// bool初始值为 false
	var d bool
	fmt.Println(d)

	// 根据值自动判断变量类型
	var e = true
	fmt.Println(e)

	// 省略 var, := 声明
	f := "hello"
	fmt.Println(f)

	// 多变量声明
	var g, h int = 1, 2
	fmt.Println(g, h)
}
