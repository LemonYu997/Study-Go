package main

import (
	"bufio"
	"fmt"
	"os"
)

// 文件写入-逐行写入
func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "Hello, World!")
	writer.Flush()
}
