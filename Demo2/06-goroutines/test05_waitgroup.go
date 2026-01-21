package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	// sync.WaitGroup 用于等待多个 Goroutine 完成。
	defer wg.Done() // Goroutine 完成时调用 Done()
	fmt.Printf("Worker %d started\n", id)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加计数器
		go worker(i, &wg)
	}

	wg.Wait() // 等待所有 Goroutine 完成
	fmt.Println("All workers done")
}
