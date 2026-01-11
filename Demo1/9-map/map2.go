package main

import "fmt"

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("key = ", key, ", value = ", value)
	}
}

func changeMap(m map[string]string) {
	// 引用传递
	m["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "New York"

	// 遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key, ", value = ", value)
	}

	fmt.Println("------------------------------------------")

	// 删除
	delete(cityMap, "Japan")
	printMap(cityMap)

	fmt.Println("------------------------------------------")
	// 修改
	cityMap["USA"] = "DC"
	printMap(cityMap)

	fmt.Println("------------------------------------------")
	// map为引用传递
	changeMap(cityMap)
	printMap(cityMap)
}
