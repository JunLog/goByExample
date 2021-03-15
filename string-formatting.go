package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p) // 打印结构体

	fmt.Printf("%+v\n", p) // 打印结构体（包含字段名）

	fmt.Printf("%#v\n", p) // 会产生该值的源码片段

	fmt.Printf("%T\n", p) // 值的类型

	fmt.Printf("%t\n", true) // 格式化布尔值

	fmt.Printf("%d\n", 123) // 格式化10进制的整数

	fmt.Printf("%b\n", 14) //格式化二进制

	fmt.Printf("%c\n", 33) //输出给定整数的对应字符

	fmt.Printf("%x\n", 456) //转 16 进制

	fmt.Printf("%f\n", 78.9) // 十进制的浮点型

	fmt.Printf("%e\n", 123400000.0) // 浮点型格式化为科学记数法e
	fmt.Printf("%E\n", 123400000.0) // 浮点型格式化为科学记数法E

	fmt.Printf("%s\n", "\"string\"") // 基本的字符串

	fmt.Printf("%q\n", "\"string\"") // 带双引号的字符串输出

	fmt.Printf("%x\n", "hex this") // hex base16 转16进制

	fmt.Printf("%p\n", &p) // 输出一个指针的值

	fmt.Printf("|%6d|%6d|\n", 12, 345) // %后数字是宽度,6个字符的宽度，空格填充，默认右对齐

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) //.后数字是精度

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // 左对齐用 -

	fmt.Printf("|%6s|%6s|\n", "foo", "b") // 字符 宽度、对齐

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // 字符 宽度、对齐

	s := fmt.Sprintf("a %s", "string") // Sprintf 是返回一个字符串，不打印
	fmt.Println(s)

	// Fprintd 输出到 io.Writers 而不是 os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
