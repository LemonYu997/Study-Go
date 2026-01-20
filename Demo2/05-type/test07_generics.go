package main

import "fmt"

/**
泛型
类型参数（Type Parameters）：允许你在函数或类型定义中使用一个或多个类型作为参数。
类型约束（Type Constraints）：指定类型参数必须满足的条件，确保在函数内部可以安全地操作这些类型。
any	约束类型参数为任何类型。	[T any]
comparable	约束类型参数为可比较的类型。	[K comparable]


// 基本语法结构
func 函数名[T 约束](参数 T) 返回值类型 {
    // 函数体
}

type 类型名[T 约束] struct {
    // 结构体字段
}
*/

// 交换两个值
func Swap[T any](a, b T) (T, T) {
	return b, a
}

// 判断切片是否包含元素
func Contains[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// 去重函数
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

// 使用示例
func main() {
	// Swap 示例
	a, b := 10, 20
	a, b = Swap(a, b)
	fmt.Printf("a=%d, b=%d\n", a, b) // 输出: a=20, b=10

	// Contains 示例
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Contains(numbers, 3)) // 输出: true

	// Unique 示例
	duplicates := []int{1, 2, 2, 3, 4, 4, 5}
	unique := Unique(duplicates)
	fmt.Println(unique) // 输出: [1 2 3 4 5]
}
