package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
递归函数通常包含两个部分：
基准条件（Base Case）：这是递归的终止条件，防止函数无限调用自身。
递归条件（Recursive Case）：这是函数调用自身的部分，用于将问题分解为更小的子问题。
*/

// 递归函数计算阶乘
func factorial(n int) int {
	// 基准条件
	if n == 0 {
		return 1
	}
	// 递归条件
	return n * factorial(n-1)
}

// 斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

// 求平方根
func sqrtRecursive(x, guess, prevGuess, epsilon float64) float64 {
	if diff := guess*guess - x; diff < epsilon && -diff < epsilon {
		return guess
	}

	newGuess := (guess + x/guess) / 2
	if newGuess == prevGuess {
		return guess
	}

	return sqrtRecursive(x, newGuess, guess, epsilon)
}

func sqrt(x float64) float64 {
	return sqrtRecursive(x, 1.0, 0.0, 1e-9)
}

// 文件目录遍历
func walkDir(dir string, indent string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		fmt.Println(indent + entry.Name())
		if entry.IsDir() {
			walkDir(filepath.Join(dir, entry.Name()), indent+"  ")
		}
	}
}

func main() {
	fmt.Println(factorial(5)) // 输出: 120

	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}

	x := 25.0
	result := sqrt(x)
	fmt.Printf("%.2f 的平方根为 %.6f\n", x, result)

	walkDir(".", "")
}
