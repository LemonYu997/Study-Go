package main

import "fmt"

// 传入动态数组
func printArray2(myArray []int) {
	// _ 表示匿名变量
	for _, value := range myArray {
		fmt.Println("value:", value)
	}

	// 动态数组为引用传递，会改变原值
	myArray[0] = 111
}

func main() {
	// 动态数组，切片 slice
	myArray := []int{1, 2, 3, 4}
	fmt.Printf("myArray type: %T\n", myArray)

	printArray2(myArray)
	// 动态数组为引用传递，会改变原值
	fmt.Println("myArray[0] = ", myArray[0])

	// 声明slice1是一个切片，默认值1，2,3, 长度为3
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice1 type: %T\n", len(slice1), slice1)
	//声明 slice2 是一个切片，但是并没有给slice分配空间
	var slice2 []int
	fmt.Printf("len = %d, slice2 type: %T\n", len(slice2), slice2)

	//声明 slice3 是一个切片，并分配3个空间
	//var slice3 []int = make([]int, 3)
	slice3 := make([]int, 3) // 简写，推荐使用这种写法
	fmt.Printf("len = %d, slice3 type: %T\n", len(slice3), slice3)

	// 判断一个slice是否为空，即没有元素
	if slice2 == nil {
		fmt.Println("slice2 is nil")
	} else {
		fmt.Println("slice2 is not nil")
	}
}
