package main

import "fmt"

func printSlice(myArray []int) {
	// _ 表示匿名变量
	for _, value := range myArray {
		fmt.Println("value:", value)
	}
}

// 切片容量追加
func main() {
	//声明切片，长度为3，容量为5
	var numbers = make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, numbers type: %T\n", len(numbers), cap(numbers), numbers)
	printSlice(numbers)

	//向numbers切片追加一个元素1，numbers len = 4, [0,0,0,1]
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, numbers type: %T\n", len(numbers), cap(numbers), numbers)
	printSlice(numbers)

	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, numbers type: %T\n", len(numbers), cap(numbers), numbers)
	printSlice(numbers)

	// 向容量 cap 已满的 slice 追加元素，此时会cap会扩容至 5 + 5 = 10
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, numbers type: %T\n", len(numbers), cap(numbers), numbers)
	printSlice(numbers)

	fmt.Println("----------------------------------")

	// 如果初始化时不指定 cap，则会和 len 一致
	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, numbers2 type: %T\n", len(numbers2), cap(numbers2), numbers2)
	printSlice(numbers2)

	// 此时再追加，len = 3 + 1 = 4, cap = 3 + 3 = 6
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, numbers2 type: %T\n", len(numbers2), cap(numbers2), numbers2)
	printSlice(numbers2)
}
