package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 用键值对的方式设置环境变量
	os.Setenv("FOO", "1")
	// 获取环境变量，没有返回为空
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	// os.Environ 列出所有的环境变量键值对
	for _, e := range os.Environ() {
		// fmt.Println(e)
		// splintN 相对于普通的 splite 可以定义最多切割成几份
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
