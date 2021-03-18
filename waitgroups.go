package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup 主要避免 goroutine 泄露

// 在每个 goroutine 启动前增加一个计数，同时在每一个 goroutine 结束时递减计数

// Add(delta int)：增加/减少若干计数
// Done：减少 1 个计数，等价于 Add(-1)
// Wait：卡住，直到计数等于 0

// 每个协程都调用这个函数，WaitGroup 必须通过指针传递
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 最后通知 WaitGroup,当前协程已经完成

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // 计数器
		go worker(i, &wg)
	}
	wg.Wait() //阻塞，直至WaitGroup计数器回复为零（即所有协程工作完成）
}
