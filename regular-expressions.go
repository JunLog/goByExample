package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// 测试一个字符串是否符合一个表达式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	// 一个 正则 的结构体（表达式）
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 匹配测试
	fmt.Println(r.MatchString("peach")) //true

	// 查找（首次）匹配的字符串
	fmt.Print("FindString	")
	fmt.Println(r.FindString("peach punch")) // peach

	// 返回首次匹配的字符串开始、结束索引
	fmt.Print("FindStringIndex	")
	fmt.Println(r.FindStringIndex("peach punch")) //[0 5]

	// 完全匹配 和 局部匹配 的字符串
	fmt.Print("FindStringSubmatch	")
	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	// 完全匹配 和 局部匹配 的字符串的开始、结束索引
	fmt.Print("FindStringSubmatchIndex	")
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) //[0 5 1 3]

	// 加多一个 All ，所有匹配项
	fmt.Print("FindAllString	")
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // 匹配次数。 [peach punch pinch]

	fmt.Print("FindAllStringSubmatchIndex	")
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", 2)) // [[0 5 1 3] [6 11 7 9]]

	fmt.Print("FindAllString	")
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	// MatchString,去掉 String ，可以不用字符串作为参数，传入一个 []byte 参数
	fmt.Println(r.Match([]byte("peach"))) // true

	// 创建 正则 表达式，使用 panic 代替返回错误（更安全）
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r) // p([a-z]+)ch

	// 替换
	fmt.Print("ReplaceAllString	")
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	in := []byte("a peach")
	// 使用 给定的函数 来转换匹配到的文本
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH

}
