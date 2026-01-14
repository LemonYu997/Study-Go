package main

import "fmt"

// 访问和修改数组
func main() {
	// 创建一个3x3的矩阵
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// 访问单个元素
	fmt.Println("第一行第二列:", matrix[0][1]) // 输出: 2
	fmt.Println("第三行第三列:", matrix[2][2]) // 输出: 9

	// 访问整行
	fmt.Println("第二行:", matrix[1]) // 输出: [4 5 6]

	// 遍历所有元素
	fmt.Println("\n遍历所有元素:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}

	// 创建一个2x2的零值矩阵
	var grid [2][2]int

	fmt.Println("修改前的矩阵:", grid)

	// 修改特定位置的元素
	grid[0][0] = 10
	grid[0][1] = 20
	grid[1][0] = 30
	grid[1][1] = 40

	fmt.Println("修改后的矩阵:", grid)

	// 批量修改一行
	grid[0] = [2]int{100, 200}
	fmt.Println("修改第一行后:", grid)
}
