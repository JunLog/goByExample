package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器，想要在未来某一刻执行一次时使用
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C // 会一直阻塞，知道定时器通道 C 发送定时器失效的值
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second) // 单纯的等待（阻塞）
}
