package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

// 状态协程
// 通过通信实现共享内存
func main() {
	var readOps uint64  // 记录读的次数
	var writeOps uint64 //记录写的次数
	// 有一个 state 共享
	// 其他协程通过 读通道 和 写通道 对 state 进行道谢
	//
	reads := make(chan readOp)   // 读通道
	writes := make(chan writeOp) // 写通道

	// 共享的 state ，一直在监听 读写两个通道
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads: // 接收 读通道 数据
				// 从 读通道 的结构体数据中的通道返回数据
				read.resp <- state[read.key]
			case write := <-writes: // 接收 写通道 数据
				// 在 state 中写数据
				state[write.key] = write.val
				// 写数据完成，在 写通道 接收到的结构体数据中的 回复通道 回复
				write.resp <- true
			}
		}
	}()

	// 启动 100 个协程，读数据
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// 构造 读通道 的数据
				read := readOp{
					key:  rand.Intn(5),   // 0~5的随机整数
					resp: make(chan int)} // 回复通道
				reads <- read                 // 将数据发送到读通道
				<-read.resp                   // 阻塞等待返回读到的数据
				atomic.AddUint64(&readOps, 1) // 原子计数，记录读数据次数
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 启动 10 个协程 写数据
	for w := 0; w < 10; w++ {
		go func() {
			for {
				// 构造 写通道 数据
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write                // 往 写通道 里发送数据
				<-write.resp                   // 阻塞等待，回复（写入完成）
				atomic.AddUint64(&writeOps, 1) // 原子计数，记录写数据次数
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps", writeOpsFinal)
}

// 基于协程的方法比基于互斥锁的方法要复杂得多
// 某些情况下它可能很有用
// 当涉及其他通道，或者管理多个同类互斥锁时，会很容易出错
