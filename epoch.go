package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	secs := now.Unix() // 时间戳（秒）
	nanos := now.UnixNano() // 时间戳（纳秒）
	fmt.Println(now)

	millis := nanos / 1000000 // 时间戳（毫秒）
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 将 时间戳 转化为时间
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(1, nanos))
}
