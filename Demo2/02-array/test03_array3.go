package main

import "fmt"

// 多维数组的常见操作
func main() {
	// 1、使用 range 遍历数组
	// 创建一个二维数组
	scores := [3][4]int{
		{85, 90, 78, 92},
		{88, 76, 95, 89},
		{92, 85, 88, 90},
	}

	fmt.Println("学生成绩表:")

	// 使用 range 遍历二维数组
	for i, row := range scores {
		fmt.Printf("学生 %d 的成绩: ", i+1)
		for j, score := range row {
			fmt.Printf("%d ", score)
			// 如果需要索引和值都使用
			_ = j // 避免未使用变量警告
		}
		fmt.Println()
	}

	// 只关心值，不关心索引
	total := 0
	count := 0
	for _, row := range scores {
		for _, score := range row {
			total += score
			count++
		}
	}
	fmt.Printf("\n平均分: %.2f\n", float64(total)/float64(count))

	fmt.Println("------------------------------------")

	// 2、获取数组长度 len
	// 创建一个不规则的多维数组
	//jagged := [3][3]int{
	//	{1, 2, 3},
	//	{4, 5},
	//	{6, 7, 8, 9}, // 注意：这里会编译错误，因为每行必须长度一致
	//}

	// 正确的示例：获取数组维度
	matrix := [4][5]int{}

	// 获取行数
	rows := len(matrix)
	fmt.Println("行数:", rows) // 输出: 4

	// 获取第一行的列数（所有行长度相同）
	cols := len(matrix[0])
	fmt.Println("列数:", cols) // 输出: 5

	// 获取总元素数
	totalElements := rows * cols
	fmt.Println("总元素数:", totalElements) // 输出: 20

	fmt.Println("------------------------------------")
	// 3、两个数组比较
	// 创建两个相同的二维数组
	a := [2][2]int{{1, 2}, {3, 4}}
	b := [2][2]int{{1, 2}, {3, 4}}
	c := [2][2]int{{1, 2}, {3, 5}}

	// 数组可以直接比较（只有当维度完全相同时）
	fmt.Println("a == b:", a == b) // 输出: true
	fmt.Println("a == c:", a == c) // 输出: false

	// 注意：不同维度的数组不能比较
	// d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(a == d)  // 编译错误：类型不匹配
}
