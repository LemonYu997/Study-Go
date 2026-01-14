package main

import (
	"fmt"
	"time"
)

// 子 goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主 goroutine
func main() {
	// 创建一个 go 程，去执行newTask()
	go newTask()

	// 主线程执行
	//i := 0
	//for {
	//	i++
	//	fmt.Printf("main Goroutine : i = %d\n", i)
	//	time.Sleep(1 * time.Second)
	//}

	// 主线程结束时所有子线程都会结束
	fmt.Println("main goroutine exit")
}
