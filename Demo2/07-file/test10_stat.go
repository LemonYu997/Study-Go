package main

import (
	"fmt"
	"log"
	"os"
)

// 文件信息与操作
func main() {
	// 检查文件是否存在
	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("文件存在")
	}

	//获取文件信息
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("文件名:", fileInfo.Name())
	fmt.Println("文件大小:", fileInfo.Size(), "字节")
	fmt.Println("权限:", fileInfo.Mode())
	fmt.Println("最后修改时间:", fileInfo.ModTime())
	fmt.Println("是目录吗:", fileInfo.IsDir())

	// 重命名和移动文件
	err = os.Rename("old.txt", "new.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("重命名成功")
}
