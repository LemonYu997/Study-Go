package main

import "fmt"

func main() {
	// 声明 map1 的 key 类型为 string ， value 类型为 string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 == nil")
	}

	// 在使用map前需要开辟空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "go"
	myMap1["three"] = "python"
	fmt.Println(myMap1)

	// 直接make声明
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "go"
	myMap2[3] = "python"
	fmt.Println(myMap2)

	// 声明时赋值
	myMap3 := map[string]string{
		"one": "java",
		"two": "go",
		// 注意最后一行也许要 ,
		"three": "python",
	}
	fmt.Println(myMap3)
}
