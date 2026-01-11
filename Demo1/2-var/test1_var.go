// 四种变量的声明方式
package main

import "fmt"

// 声明全局变量，方法1、2、3可以
var gA int = 100
var gB = 200

// 方法4 全局变量无法声明
//gC := 300

func main() {
	// 方法1：声明一个变量，默认0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法2：声明一个变量，初始化值
	var b int = 100
	fmt.Println("b=", b)
	fmt.Printf("type of b = %T\n", b)
	var bb string = "hello"
	fmt.Println("bb=", bb)
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	// 方法3：初始化时省略数据类型，自动匹配
	var c = 100
	fmt.Println("c=", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "hello"
	fmt.Println("cc=", cc)
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	// 方法4：省略 var，自动匹配 := 自动初始化，不支持全局变量声明方式
	d := 100
	fmt.Println("d=", d)
	fmt.Printf("type of d = %T\n", d)

	e := "hello"
	fmt.Println("e=", e)
	fmt.Printf("e = %s, type of e = %T\n", e, e)

	g := 1.23
	fmt.Println("g=", g)
	fmt.Printf("type of g = %T\n", g)

	// 全局变量区别
	fmt.Println("gA = ", gA)
	fmt.Println("gB = ", gB)
	//fmt.Println("gC = ", gC)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx=", xx, "yy=", yy)
	var kk, ll = 100, "hello"
	fmt.Println("kk=", kk, "ll=", ll)

	// 多行变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv=", vv)
	fmt.Println("jj=", jj)
}
