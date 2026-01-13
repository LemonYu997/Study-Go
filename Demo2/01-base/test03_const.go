package main

import "fmt"

const (
	/// 自动修改常量
	a = iota //0
	b        //1
	c        //2
	d = "ha" //独立值，iota += 1
	e        //"ha"   iota += 1
	f = 100  //iota +=1
	g        //100  iota +=1
	h = iota //7,恢复计数
	i        //8
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	// 多变量多赋值
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Println(area)
	fmt.Println(a, b, c)
}
