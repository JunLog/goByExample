package main

import "fmt"

func main() {
	message := make(chan string) // 创建一个新的通道

	go func() {
		message <- "ping"
		// channel <- 信息
		// 发送一个消息到通道（channel）
	}()

	msg := <-message // 默认通道的发送和接收是阻塞的
	// 接收 <- channel
	fmt.Println(msg) // 可以一直等到有信息才打印
}
