package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	// done 通道用来通知其他协程已经完成工作
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // 发送信息
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done // 接受信息（阻塞），因此不得到信息，程序不会停止
}
