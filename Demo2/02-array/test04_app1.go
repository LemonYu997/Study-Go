package main

import "fmt"

// 数组应用1 游戏棋盘（井字棋）
func main() {
	// 初始化一个3x3的井字棋棋盘
	var board [3][3]string

	// 初始化为空
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}

	// 模拟下棋
	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"

	// 打印棋盘
	fmt.Println("井字棋棋盘:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" %s ", board[i][j])
			if j < 2 {
				fmt.Printf("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}
}
