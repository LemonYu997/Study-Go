package main

import (
	"fmt"
	"math"
)

/*
*
接口
隐式实现：

	Go 中没有关键字显式声明某个类型实现了某个接口。
	只要一个类型实现了接口要求的所有方法，该类型就自动被认为实现了该接口。

接口类型变量：

	接口变量可以存储实现该接口的任意值。
	接口变量实际上包含了两个部分：

零值接口：

	接口的零值是 nil。
	一个未初始化的接口变量其值为 nil，且不包含任何动态类型或值。

空接口：

	定义为 interface{}，可以表示任何类型。
*/

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义一个结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	c := Circle{Radius: 5}
	var s Shape = c // 接口变量可以存储实现了接口的类型
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())

	// 动态值和动态类型
	var i interface{} = 42
	fmt.Printf("Dynamic type: %T, Dynamic value: %v\n", i, i)

	// 接口的零值
	var i2 interface{}
	fmt.Println(i2 == nil) // 输出：true

	// 类型选择
	printType(42)
	printType("hello")
	printType(3.14)
	printType([]int{1, 2, 3})
}

func printType(val interface{}) {
	// 类型选择 根据接口变量的具体类型执行不同的逻辑
	switch v := val.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float:", v)
	default:
		fmt.Println("Unknown type")
	}
}
