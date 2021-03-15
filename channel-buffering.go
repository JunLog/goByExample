package main

import "fmt"

func main() {
	message := make(chan string, 2)
	// 默认channel(通道)是没有缓冲的，即 通道里有信息未读取的话，不能再往通道里发东西
	// 有 缓冲 ，即可在未读取通道信息的情况下，缓存多次发送的信息
	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}

// 通道的 标识符 是 <-
// 读取 与 发送 的区别在于，通道在哪边
// 看 最左边 的变量，确定行为（是 读 还是 发）
// 被箭头指向通道（channel 在左），就很明显是，发送信息到通道里来
// 箭头是由通道往外指（channe 在右），就是通道发送信息出去
