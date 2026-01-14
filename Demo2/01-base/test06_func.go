package main

import (
	"fmt"
	"math"
)

// 函数名(形参) 返回类型
func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 多个返回值
func swap(x, y string) (string, string) {
	return y, x
}

// 函数的传递方式-值传递
func swap2(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp

	return temp
}

func swap3(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

// 函数闭包案例-返回另一个函数
func getSequence() func() int {
	i := 0
	// 返回一个匿名函数
	// 匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
	return func() int {
		i += 1
		return i
	}
}

/* 定义结构体 */
type Circle struct {
	radius float64
}

// 该 method 属于 Circle 类型对象中的方法
// 注意和普通函数的区别，就是函数名前声明了方法接收者（这里是Circle对象）
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func main() {
	a := 100
	b := 200
	var ret int
	ret = max(a, b)
	fmt.Println("The max = ", ret)

	c, d := swap("hello", "world")
	fmt.Println(c, d)

	// 函数调用-值传递
	fmt.Printf("交换前 a 的值为 : %d\n", a) //100
	fmt.Printf("交换前 b 的值为 : %d\n", b) //200
	swap2(a, b)
	fmt.Printf("swap2 交换后 a 的值 : %d\n", a) //100
	fmt.Printf("swap2 交换后 b 的值 : %d\n", b) //200

	// 函数调用-引用传递
	swap3(&a, &b)
	fmt.Printf("swap3 交换后 a 的值 : %d\n", a) //200
	fmt.Printf("swap3 交换后 b 的值 : %d\n", b) //100

	// 将函数本身作为参数
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber()) //1
	fmt.Println(nextNumber()) //2
	fmt.Println(nextNumber()) //3

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1()) //1
	fmt.Println(nextNumber1()) //2

	// 定义一个匿名函数并将其赋值给变量add
	add := func(a, b int) int {
		return a + b
	}

	// 调用匿名函数
	result := add(3, 5)
	fmt.Println("3 + 5 =", result) //3 + 5 = 8

	// 在函数内部使用匿名函数
	multiply := func(x, y int) int {
		return x * y
	}

	product := multiply(4, 6)
	fmt.Println("4 * 6 =", product) //4 * 6 = 24

	// 将匿名函数作为参数传递给其他函数
	calculate := func(operation func(int, int) int, x, y int) int {
		return operation(x, y)
	}

	sum := calculate(add, 2, 8)
	fmt.Println("2 + 8 =", sum) //2 + 8 = 10

	// 也可以直接在函数调用中定义匿名函数
	difference := calculate(func(a, b int) int {
		return a - b
	}, 10, 4)
	fmt.Println("10 - 4 =", difference) //10 - 4 = 6

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}
