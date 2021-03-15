package main

import "fmt"

func intSeq() func() int { // 返回函数内定义的匿名函数
	i := 0
	return func() int { //匿名函数
		i++
		return i
	}
}
func main() {
	nextInt := intSeq() //调用闭包函数，里面的 i 相当于是全局变量

	fmt.Println(nextInt()) // 调用闭包内部的匿名函数时，i 有值
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() // 调用的闭包函数，新建一个全局
	fmt.Println(newInts())

	fmt.Println(nextInt())
}

// 闭包因为将变量都保存在内存中了，会加大内存消耗，严重时可能会发生内存泄露
