package main

import (
	"log"
	"os"
)

// 文件创建
func main() {
	// 创建文件，如果文件已存在会被截断（清空）
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // 确保文件关闭

	log.Println("文件创建成功")
}
