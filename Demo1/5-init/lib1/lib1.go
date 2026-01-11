package lib1

import "fmt"

// 当前lib1包提供的API 首字母必须大写
func Lib1Test() {
	fmt.Println("lib1 Test")
}

func init() {
	fmt.Println("lib1 init")
}
