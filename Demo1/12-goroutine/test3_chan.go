package main

import "fmt"

func main() {
	// 定义一个 channel，用来在线程之间通信
	c := make(chan int)

	go func() {
		defer fmt.Printf("goroutine exit\n")
		fmt.Println("goroutine 正在执行..")

		// 将 666 发送给c
		c <- 666
	}()

	// 从 c 中接收数据并赋值给 num
	// 因为存在接收和发送，主线程会阻塞，直到接收到 c 中的数据
	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("main goroutine exit")
}
