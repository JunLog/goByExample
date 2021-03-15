package main

import "fmt"

// 通道在作为函数的参数时，可以指定 只读 或 只写（提高安全性）
// chan<- 只写
// <-chan 只读
// 还是看 通道 在哪边
// 看左边是变量（只写）还是标识符（只读）

func ping(pings chan<- string, msg string) { // 只写
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// pings 只读，pongs 只写
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
