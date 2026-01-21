package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 创建单个目录
	err := os.Mkdir("newdir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// 递归创建多级目录
	err = os.MkdirAll("path/to/newdir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// 读取目录内容
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		info, _ := entry.Info()
		fmt.Printf("%-20s %8d %v\n",
			entry.Name(),
			info.Size(),
			info.ModTime().Format("2006-01-02 15:04:05"))
	}

	// 删除空目录
	err = os.Remove("emptydir")
	if err != nil {
		log.Fatal(err)
	}

	// 递归删除目录及其内容
	err = os.RemoveAll("path/to/dir")
	if err != nil {
		log.Fatal(err)
	}
}
