package main

import "fmt"

// 主动关闭 chan
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		// close可以关闭一个 chan，关闭后无法再发送数据
		// 但是关闭后仍可以从 c 中接收数据
		// 如果chan 类型为 nil,无论收发都会被阻塞
		close(c)
	}()

	//for {
	//	// ok如果为 true 表示 channel 没有关闭，如果为false表示 channel已经关闭
	//	if data, ok := <-c; ok {
	//		fmt.Println(data)
	//	} else {
	//		break
	//	}
	//}

	// 使用 range 简写上述代码，可以从 c 中获取数据
	// range可以迭代不断操作 channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main goroutine exit")
}
