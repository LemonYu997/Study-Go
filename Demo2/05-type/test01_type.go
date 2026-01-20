package main

import (
	"fmt"
	"strconv"
)

// 基础类型转换
func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	// int 转为 float
	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)

	// string 转 int
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	}

	// int 转为字符串
	num = 123
	str = strconv.Itoa(num)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num, str)

	// string 转 float
	str = "3.14"
	num2, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str, num2)
	}

	// float 转 string
	num3 := 3.14
	str3 := strconv.FormatFloat(num3, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num3, str3)
}
