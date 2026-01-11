package lib2

import "fmt"

// 当前lib1包提供的API 首字母必须大写
func Lib2Test() {
	fmt.Println("lib2 Test")
}

func init() {
	fmt.Println("lib2 init")
}
