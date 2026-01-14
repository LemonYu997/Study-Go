package main

import "fmt"

func main() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j, k int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	fmt.Println("-------------------------------------")

	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	// 不确定数组长度，编译器根据初始化元素数自动判断
	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	fmt.Println("-------------------------------------")
	// 二维数组
	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])

	// 初始化二维数组
	// 创建二维数组
	sites := [2][2]string{}

	// 向二维数组添加元素
	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"

	// 显示结果
	fmt.Println(sites)

	// 访问二维数组
	/* 数组 - 5 行 2 列*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

	//创建各个维度元素数量不一致的多维数组
	// 创建空的二维数组
	animals := [][]string{}

	// 创建三一维数组，各数组长度不同
	rowA1 := []string{"fish", "shark", "eel"}
	rowA2 := []string{"bird"}
	rowA3 := []string{"lizard", "salamander"}

	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, rowA1)
	animals = append(animals, rowA2)
	animals = append(animals, rowA3)

	// 循环输出
	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}

	fmt.Println("-------------------------------------")
	// 三维数组

	// 声明一个2x3x4的三维数组
	// 可以理解为：2个平面，每个平面有3行4列
	var cube [2][3][4]int

	// 初始化三维数组
	cube = [2][3][4]int{
		{ // 第一个平面
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		{ // 第二个平面
			{13, 14, 15, 16},
			{17, 18, 19, 20},
			{21, 22, 23, 24},
		},
	}

	// 访问三维数组元素
	fmt.Println("cube[0][1][2] =", cube[0][1][2]) // 输出: 7
	fmt.Println("cube[1][2][3] =", cube[1][2][3]) // 输出: 24

	// 遍历三维数组
	fmt.Println("\n三维数组内容:")
	for i := 0; i < 2; i++ {
		fmt.Printf("平面 %d:\n", i)
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				fmt.Printf("%3d ", cube[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
