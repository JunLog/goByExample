package main

import (
	"fmt"
	"sync"
	"time"
)

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
