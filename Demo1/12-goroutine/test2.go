package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 用 go 创建承载一个形参为空，返回值为空的匿名函数
	go func() {
		defer fmt.Println("A.defer")

		// 子函数
		func() {
			defer fmt.Println("B.defer")

			// 如果想在执行过程中结束掉整个 goroutine ，需要执行以下方法
			runtime.Goexit()

			fmt.Println("B")
			// 注意结尾要加()表示调用该匿名函数
		}()

		fmt.Println("A")
		// 注意结尾要加()表示调用该匿名函数
	}()

	// 调用带形参的匿名函数
	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		// 由于这里是以子线程的形式执行的，因此这里的返回值在主线程中是无法接收的
		return true

		// 注意结尾要加()表示调用该匿名函数，这里需要传入对应的参数
	}(10, 20)

	// 主线程死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
