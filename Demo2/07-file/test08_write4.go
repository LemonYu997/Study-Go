package main

import (
	"fmt"
	"os"
)

// 文件的追加写入
func main() {
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 追加写入
	if _, err := file.WriteString("Appended text\n"); err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	fmt.Println("Text appended successfully!")
}
