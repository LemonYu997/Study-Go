package main

import (
	"fmt"
	"log"
	"os"
)

// 文件写入
func main() {
	// 方式1：直接写入字符串
	file, err := os.Create("write1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("直接写入字符串\n")

	// 方式2：写入字节切片
	data := []byte("写入字节切片\n")
	file.Write(data)

	// 方式3：使用fmt.Fprintf格式化写入
	fmt.Fprintf(file, "格式化写入: %d\n", 123)
}
