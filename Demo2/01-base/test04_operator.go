package main

import "fmt"

/**
运算符：
- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
- 赋值运算符
- 其他运算符
*/

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c) // 31
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c) // 11
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c) // 210
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c) // 2
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c) // 1
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a) // 22
	a = 21                            // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a) // 20

	fmt.Println("-----------------------------------------------------")
	// 关系运算符
	if a == b {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if a < b {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if a > b {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第四行 - a 小于等于 b\n")
	}
	if b >= a {
		fmt.Printf("第五行 - b 大于等于 a\n")
	}

	fmt.Println("-----------------------------------------------------")
	// 逻辑运算符
	var d bool = true
	var e bool = false
	if d && e {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if d || e {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 d 和 e 的值 */
	d = false
	e = true
	if d && d {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(d && e) {
		fmt.Printf("第四行 - 条件为 true\n")
	}

	fmt.Println("-----------------------------------------------------")

	// 位运算符
	var f uint = 60 /* 60 = 0011 1100 */
	var g uint = 13 /* 13 = 0000 1101 */
	var h uint = 0

	h = g & f /* 12 = 0000 1100 */
	fmt.Printf("第一行 - h 的值为 %d\n", h)

	h = g | f /* 61 = 0011 1101 */
	fmt.Printf("第二行 - h 的值为 %d\n", h)

	h = g ^ f /* 49 = 0011 0001 */
	fmt.Printf("第三行 - h 的值为 %d\n", h)

	h = g << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - h 的值为 %d\n", h)

	h = g >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - h 的值为 %d\n", h)

	fmt.Println("-----------------------------------------------------")

	a = 21
	c = 0
	c = a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c) //21

	c += a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c) //42

	c -= a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c) //21

	c *= a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c) //441

	c /= a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c) //21

	c = 200

	c <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c) //800

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c) //200

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c) //0

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c) //2

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c) //2

	fmt.Println("-----------------------------------------------------")

	// 其他运算符 &地址 *指针
	var i int = 4
	var j int32
	var k float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - i 变量类型为 = %T\n", i)
	fmt.Printf("第 2 行 - j 变量类型为 = %T\n", j)
	fmt.Printf("第 3 行 - k 变量类型为 = %T\n", k)

	/*  & 和 * 运算符实例 */
	ptr = &i /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("i 的值为  %d\n", i)
	fmt.Printf("*ptr 为 %d\n", *ptr)
	fmt.Println("-----------------------------------------------------")

	// 运算符优先级
	var l int = 20
	var m int = 10
	var n int = 15
	var o int = 5
	var p int

	p = (l + m) * n / o // ( 30 * 15 ) / 5
	fmt.Printf("(l + m) * n / o 的值为 : %d\n", p)

	p = ((l + m) * n) / o // (30 * 15 ) / 5
	fmt.Printf("((l + m * n) / o 的值为  : %d\n", p)

	p = (l + m) * (n / o) // (30) * (15/5)
	fmt.Printf("(l + m) * (n / o) 的值为  : %d\n", p)

	p = l + (m*n)/o //  20 + (150/5)
	fmt.Printf("l + (m * n) / o 的值为  : %d\n", p)
}
