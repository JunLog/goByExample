package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}
// var 可声明 1个 或 多个 变量
// 变量初始化没赋值，那么初始化会是 零值
// := 是简化的声明并初始化(只能在函数里使用)
