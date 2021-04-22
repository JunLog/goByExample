package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// jobs、results 通道参数，<-chan 只读， chan<- 只写
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动 3 个 worker 工作池（请了 3 个员工在待命）
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j // 分配任务
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		<-results // 收集返回值，确保所有协程已完成
	}
}
