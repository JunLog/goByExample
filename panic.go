package main

import "os"

func main() {
	// panic 意味着有些出乎意料的错误发生
	panic("小问题") // 触发 panic

	_, err := os.Create("./tmp/file")
	// 函数返回我们不知道如何处理（或不想处理）的错误值时
	if err != nil {
		panic(err)
		// 输出 一个错误消息 和 协程跟踪信息
		// 非零状态 退出
	}
}

// 这里创建文件失败，貌似是因为：
// 1. 前面有一个 panic （有点奇怪，这里的机制不太懂）
// 2. 路径不存在（没有 tmp 这个文件夹）
