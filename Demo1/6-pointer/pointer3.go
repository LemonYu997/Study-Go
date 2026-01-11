package main

import "fmt"

// 值传递时不能真正改变原先两个变量的值
/*func swap1(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}*/

// 使用指针，引用传递时，可以真正改变
func swap2(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a int = 10
	var b int = 20
	fmt.Println("swap before , a = ", a, "b = ", b)

	//swap1(a, b)	// 10, 20
	swap2(&a, &b) // 20m, 10

	fmt.Println("swap after , a = ", a, "b = ", b)

	// 两个地址相同
	var p *int
	p = &a
	fmt.Println("p = ", p)
	fmt.Println("&a = ", &a)

	// 二级指针 pp 和 &p 地址一样
	var pp **int
	pp = &p
	fmt.Println("pp = ", pp)
	fmt.Println("&p = ", &p)
}
