package main

import "fmt"

// 类型断言（Type Assertion）是一种用于检查接口值的实际类型的机制。
// value, ok := interfaceValue.(Type)
func main() {
	var i interface{} = "Hello, Go!"

	// 尝试将 i 断言为 string 类型
	s, ok := i.(string)
	if ok {
		fmt.Println("断言成功:", s)
	} else {
		fmt.Println("断言失败")
	}

	// 尝试将 i 断言为 int 类型
	n, ok := i.(int)
	if ok {
		fmt.Println("断言成功:", n)
	} else {
		fmt.Println("断言失败")
	}

	// 可以不返回布尔值，而是直接在断言失败时引发 panic。
	// 直接断言为 string 类型
	s = i.(string)
	fmt.Println("断言成功:", s)

	// 直接断言为 int 类型（会引发 panic）
	n = i.(int)
	fmt.Println("断言成功:", n)

	fmt.Println("----------------------------------------")

	// 断言的常见用法：
	// 处理多种类型的接口值
	switch v := i.(type) {
	case int:
		fmt.Println("这是一个整数:", v)
	case string:
		fmt.Println("这是一个字符串:", v)
	default:
		fmt.Println("未知类型")
	}
	// 从接口中提取具体类型
	if s, ok := i.(string); ok {
		fmt.Println("处理字符串:", s)
	} else if n, ok := i.(int); ok {
		fmt.Println("处理整数:", n)
	} else {
		fmt.Println("无法处理的类型")
	}
}
