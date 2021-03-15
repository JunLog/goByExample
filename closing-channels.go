package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			// 通道接收两个值，第一个是接收的数据，第二个值：通道关闭，且通道中所有值已经接收完毕，会是false，否则是true
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true //信道同步，通知 main() 协程
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // 关闭通道，不能再向通道发送数据
	fmt.Println("sent all jobs")
	// 信道同步
	<-done // 接受信息（阻塞），因此不得到信息，程序不会停止
}
