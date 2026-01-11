package main

import "fmt"

// 数组作为函数传参时，必须严格传入对应的大小数组
func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
	// 值拷贝，不会修改原值
	myArray[0] = 111
}

func main() {
	// 固定长度数组声明
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	// 固定遍历
	for i := 0; i < 10; i++ {
		fmt.Println(myArray1[i])
	}

	// 数组遍历，返回两个参数索引和值
	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 type: %T\n", myArray1)
	fmt.Printf("myArray2 type: %T\n", myArray2)
	fmt.Printf("myArray3 type: %T\n", myArray3)

	// 数组作为函数传参时，必须严格传入对应的大小数组
	//printArray(myArray1)
	printArray(myArray3)                       //值拷贝，不会修改原值
	fmt.Println("myArray3[0] = ", myArray3[0]) //11

}
