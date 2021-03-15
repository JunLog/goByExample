package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)
	// limiter 通道每 200ms 接收一个值。 这是我们任务速率限制的调度器。

	for req := range requests {
		<-limiter //阻塞
		fmt.Println("request", req, time.Now())
	}
	fmt.Println("///////////")
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now() // 填充通道，立刻执行
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			// 每 200 ms 传进
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5) // 填充请求
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
	// 前 3 个直接处理
	// 后两个以大约 200ms 速度处理
}
