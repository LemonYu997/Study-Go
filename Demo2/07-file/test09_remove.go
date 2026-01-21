package main

import (
	"fmt"
	"os"
)

// 文件删除
func main() {
	err := os.Remove("output.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File deleted successfully!")
}
