package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // 通道非空还是可以关闭的

	for elem := range queue {
		// 通道里的值也是可以遍历
		fmt.Println(elem)
	}
}
