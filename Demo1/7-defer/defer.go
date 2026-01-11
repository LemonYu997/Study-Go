package main

import "fmt"

func func1() {
	fmt.Println("func1")
}

func func2() {
	fmt.Println("func2")
}

func func3() {
	fmt.Println("func3")
}

func deferFunc() int {
	fmt.Println("deferFunc")
	return 0
}

func returnFunc() int {
	fmt.Println("returnFunc")
	return 1
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// 写入 defer 关键字，在函数结束前执行的函数
	defer fmt.Println("main end")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")

	// 会按栈的顺序执行 先入后出
	defer func1()
	defer func2()
	defer func3()
	// 实际顺序： func3 -> func2 -> func1 -> main end

	// 当同时有 先 return 再 defer，实际开发不推荐这样写，不易读
	returnAndDefer()
}
