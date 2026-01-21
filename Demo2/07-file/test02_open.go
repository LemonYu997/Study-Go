package main

import (
	"fmt"
	"os"
)

// 文件的打开与关闭
func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully!")
}
