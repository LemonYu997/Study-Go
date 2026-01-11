package main

import "fmt"

// 函数名(形参)返回值类型
func foo1(a string, b int) int {
	fmt.Println("----------foo1----------")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	c := 100
	return c
}

// 多个返回值 匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("----------foo2----------")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	return 666, 777
}

// 多个返回值 带形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----------foo3----------")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	// 返回形参会按默认值赋值
	// r1 , r2 作用域空间为foo3 整个函数体{}空间
	fmt.Println("r1=", r1) //0
	fmt.Println("r2=", r2) //0
	r1 = 1000
	r2 = 2000
	// 自动返回对应的形参名称
	return
}

// 返回时多个形参同一类型
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("----------foo4----------")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	r1 = 3000
	r2 = 4000
	// 自动返回对应的形参名称
	return
}

func main() {
	c := foo1("hello", 100)

	fmt.Println("c = ", c)

	ret1, ret2 := foo2("haha", 200)
	fmt.Println("ret1, ret2 = ", ret1, ret2)

	ret1, ret2 = foo3("hehe", 300)
	fmt.Println("ret1, ret2 = ", ret1, ret2)

	ret1, ret2 = foo4("heihei", 400)
	fmt.Println("ret1, ret2 = ", ret1, ret2)
}
