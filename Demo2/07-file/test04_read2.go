package main

import (
	"fmt"
	"os"
)

// 文件的读取-一次性读取整个文件
func main() {
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(content))
}
