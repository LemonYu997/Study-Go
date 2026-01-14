package main

import "fmt"

func main() {
	// 创建带缓存的 channel
	c := make(chan int, 3)

	// len(c) 当前 chan 元素数量, cap(c)当前 chan 容量
	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子go程正在运行：len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()

	for i := 0; i < 3; i++ {
		// 从 c 中接收数据，并赋值给 num
		num := <-c
		fmt.Println("num = ", num)
		fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))
	}

	fmt.Println("main goroutine exit")
}
