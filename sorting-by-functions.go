package main

import (
	"fmt"
	"sort"
)

// 自定义排序
// 按照字符串的长度进行排序

type byLength []string // 创建类型

// 实现 sort.Interface （Len、Swap、Less）接口
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

// 类似的，参照这个例子，创建一个自定义类型， 为它实现 Interface 接口的三个方法， 然后对这个自定义类型的切片调用 sort.Sort 方法， 我们就可以通过任意函数对 Go 切片进行排序了。
