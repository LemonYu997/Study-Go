package main

import (
	"fmt"
	"os"
)

// 文件写入-一次性写入
func main() {
	content := []byte("Hello, World!")
	err := os.WriteFile("output.txt", content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File written successfully!")
}
