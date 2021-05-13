package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 互斥锁 ，在协程间安全的访问数据
func main() {
	var state = make(map[int]int) // 协程操作的数据
	// map[键类型]值类型

	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {
		// 创建 100 个 协程
		go func() {
			total := 0
			for {
				key := rand.Intn(5) //随机取一个[0,5) 整数
				mutex.Lock()                  // 上锁
				total += state[key]           // 上锁期间保证独占访问
				mutex.Unlock()                // 开锁
				atomic.AddUint64(&readOps, 1) // 读取次数 +1

				time.Sleep(time.Millisecond)
				// 每 1ms 读取一次
			}
		}()
	}
	for w := 0; w < 10; w++ {
		// 创建 10 个协程
		go func() {
			for {
				key := rand.Intn(5) // 返回不大于5的非负整数
				val := rand.Intn(100)
				mutex.Lock()                   //上锁
				state[key] = val               //赋值
				mutex.Unlock()                 //解锁
				atomic.AddUint64(&writeOps, 1) //写入次数 +1
				time.Sleep(time.Millisecond)
				// 每 1ms 写入一次
			}
		}()
	}
	time.Sleep(time.Second) // 阻塞 1s

	// 获取并输出最终操作计数
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state", state)
	mutex.Unlock()
}
