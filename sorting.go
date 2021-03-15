package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	// 切片 定义是不指定长度，指定长度就是数组
	sort.Strings(strs) // sort 方法是直接改变给定切片的
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints) // 升序
	fmt.Println("Ints:	", ints)

	s := sort.IntsAreSorted(ints) // 检查是否 升序
	fmt.Println("Sorted:	", s)
}
