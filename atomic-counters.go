package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64

	var wg sync.WaitGroup

	// 启动 50 个协程
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1) // 原子计数
				// ops++ // 不使用这种，协程干扰
			}
			wg.Done()
		}()
	}
	wg.Wait() // 等待所有协程完成
	fmt.Println("ops", ops)
}
