package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now() // 获取当前时间
	// 后面的 m=+*.***** 代表的是 Monotonic Clocks,自进程启动以来的秒数
	// https://zhuanlan.zhihu.com/p/47754783
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC) // 构建时间
	p(then)

	// 提取出时间的各个部分
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond()) // 纳秒
	p(then.Location())

	p(then.Weekday())

	// 比较时间
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then) // 两个时间点的间隔时间
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))  // 时间向后移动
	p(then.Add(-diff)) // 时间向前移动

}
