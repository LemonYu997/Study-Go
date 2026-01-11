package main

import "fmt"

// const 定义枚举类型
//const (
//	// 可以在 const() 添加一个关键字 iota，每行的iota累加1，第一行默认 0
//	BEIJING  = iota // 0
//	SHANGHAI        // 1
//	SHENZHEN        // 2
//)

const (
	// 可以在 const() 添加一个关键字 iota，每行的iota累加1，第一行默认 0
	BEIJING  = 10 * iota // 0
	SHANGHAI             // 10
	SHENZHEN             // 20
)

const (
	a, b = iota + 1, iota + 2 //第0行，0 + 1 = 1, 0 + 2 = 2
	c, d                      //第1行，1 + 1 = 2, 1 + 2 = 3
	e, f                      //第2行，2 + 1 = 3, 2 + 2 = 4

	g, h = iota * 2, iota * 3 //第3行，3 * 2 = 6, 3 * 3 = 9
	i, k                      //第4行，4 * 2 = 8, 4 * 3 = 12
)

func main() {
	// 常量
	const length int = 10

	fmt.Println("length=", length)

	// 常量无法修改
	//length = 10

	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("SHENZHEN=", SHENZHEN)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)
	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)
}
