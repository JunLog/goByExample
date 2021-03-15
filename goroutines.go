package main

import (
	"fmt"
	"time"
)

func f(from string) {
	// go 协程，并发（如果 i 的判断条件太小，看不出）
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine1")
	go f("goroutine2")
	go f("goroutine3")
	go f("goroutine4")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
