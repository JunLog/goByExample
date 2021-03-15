package main

import (
	"fmt"
	"os"
)

func main() {
	// 当使用 os.Exit 时 defer 将不会被执行
	defer fmt.Println("!")

	os.Exit(3)
	// exit code int [0,125]
	// 0 表示成功，非 0 表示出错
	// 非 0 ，程序会立刻终止，并且 defer 的函数不会被执行
}
