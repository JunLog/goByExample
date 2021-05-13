package main

import (
	"fmt"
	"strings"
)

// 根据 go-lint 的提示，函数注释前 加一个函数名

//Index 返回字符串 t 在 vs 中第一次出现的索引，没有返回 -1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//Include 字符串 t 存在 切片 vs 中就返回 true
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

//Any 切片 vs 中任意一个字符满足条件 f() 就返回 true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//All vs 全满足条件 f() 返回 true
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//Filter 返回满足 f() 的新切片
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//Map() 返回一个 原切片经过 f() 映射的 新切片
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"桃子", "苹果", "梨", "李子"}
	var strs2 = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "梨")) // 返回索引号

	fmt.Println(Include(strs, "葡萄")) // 是否包含

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasSuffix(v, "子") // 是否有 这个 后缀
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasSuffix(v, "子") // 前缀是 hasPrefix()
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "子") // 返回包含该字符的切片
	}))

	fmt.Println(Map(strs2, strings.ToUpper)) // 返回经过处理的新切片
}
