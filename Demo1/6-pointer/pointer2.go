package main

import "fmt"

func changeValue2(p *int) {
	*p = 10
	fmt.Println("*p = ", *p) // 10
}

func main() {
	var a int = 1
	// 使用指针，是引用传递，会改变对应的值
	changeValue2(&a)
	fmt.Println("a = ", a) // 10
}
