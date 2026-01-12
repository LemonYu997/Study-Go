package main

import "fmt"

// 万能接口 空接口
func myFunc(arg interface{}) {
	fmt.Println("hello world")
	fmt.Println(arg)

	// interface{} 区分此时引用的数据类型到底是什么
	// 使用类型断言机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not a string")
	} else {
		fmt.Println("arg is a string, value = ", value)
	}

	fmt.Println("--------------------------------")
}

type Bookk struct {
	author string
}

func main() {
	bookk := Bookk{"Golang"}
	myFunc(bookk)
	myFunc(bookk.author)
	myFunc(100)
}
