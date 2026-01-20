package main

import "fmt"

// 定义一个接口 Writer
type Writer interface {
	Write([]byte) (int, error)
}

// 实现 Writer 接口的结构体 StringWriter
type StringWriter struct {
	str string
}

// 实现 Write 方法
func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

// 空接口类型作为函数入参，表示万能类型
func printValue(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

// 接口类型转换 类型断言和类型转换
func main() {
	var i interface{} = "Hello, World"
	// 类型断言 value.(type) 用于将接口类型转换为指定类型
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}

	// 类型转换 T(value) 将一个接口类型的值转换为另一个接口类型
	// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
	var w Writer = &StringWriter{}

	// 将 Writer 接口类型转换为 StringWriter 类型
	sw := w.(*StringWriter)

	// 修改 StringWriter 的字段
	sw.str = "Hello, World"

	// 打印 StringWriter 的字段值
	fmt.Println(sw.str)

	// 空接口类型作为函数入参，表示万能类型
	printValue(42)
	printValue("hello")
	printValue(3.14)
}
