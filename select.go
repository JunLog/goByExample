package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select { // 同时等待多个通道操作
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	elapsed := time.Since(t)

	fmt.Println("app elapsed:", elapsed)
	// 并发执行的程序，一共 2 秒左右
}
