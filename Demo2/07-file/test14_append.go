package main

import (
	"log"
	"os"
)

// 文件追加
func main() {
	file, err := os.OpenFile("log.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString("新的日志内容\n"); err != nil {
		log.Fatal(err)
	}
}
