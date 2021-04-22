package main

import "fmt"

func main() {
	messages := make(chan string,1)
	singals := make(chan bool) //常规的 通道 ，发送和接收数据是阻塞的

	select {
	case msg := <-messages: //非阻塞接收
		fmt.Println("received message", msg)
	default: //使用一个 default 子句的 select 来实现 非阻塞 发送、接收数据
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg: // 非阻塞 发送数据 方法
	// messages 通道默认是 无缓冲
	// 只有 <-messages 接收准备好的情况下，该通道发送才能执行
	// 上一个 select 并没有在准备接收
	// 可以在定义时定义为有缓冲的
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-singals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
