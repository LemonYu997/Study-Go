package main

import (
	"fmt"
	"log"
	"os"
)

// 临时文件和目录
func main() {
	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // 清理

	fmt.Println("临时文件:", tmpFile.Name())

	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "example-*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir) // 清理

	fmt.Println("临时目录:", tmpDir)
}
