package main

import (
	"flag"
	"fmt"
)

func main() {
	// 基本的标志声明 支持：字符串 、 整数 、 布尔值
	wordPtr := flag.String("word", "foo", "a string") // (标志，默认值，描述)
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool") //带上这个标志就是 true
	// 这里的都是指针变量

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	// 已有的参数也可用来声明标志，需要用到指针

	flag.Parse() // 所有标志声明完成后，调用来进行命令解析

	fmt.Println("word:", *wordPtr) // 对指针的解引用
	fmt.Println("numb:", *numbPtr) // 对指针的解引用
	fmt.Println("fork:", *boolPtr) // 对指针的解引用
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args()) // 标志位后面的参数（位置参数）
}

// go build command-line-flags.go

// 基本使用
// command-line-flags -word=opt -numb=7 -fork=7 -fork -svar=flag

// 省略一个标志标志参数的话，默认值
// command-line-flags -word=opt

//  位置参数(不带标志位的)
// command-line-flags -word=opt a1 a2 a3

// flag 包需要所有的标志出现位置参数之前（否则，这个标志将会被解析为位置参数）
// 最后的 -numb=7 被当做是位置参数
// command-line-flags.exe -word=opt a1 a2 a3 -numb=7

// 使用 -h 或者 --help 标志来得到自动生成的这个命令行程序的帮助文本
// command-line-flags.exe -h

// 用了一个没有定义的标志的话会报错，而已显示帮助文本
