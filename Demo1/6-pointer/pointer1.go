package main

import "fmt"

func changeValue1(p int) {
	p = 10
	fmt.Println("p = ", p) // 10
}

// 不使用指针时
func main() {
	var a int = 1
	// 不使用指针时，这里是值传递，会创建一个新地址存储变量 p
	changeValue1(a)
	fmt.Println("a = ", a) // 1
}
