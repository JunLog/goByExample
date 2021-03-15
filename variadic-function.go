package main

import "fmt"

func sum(nums ...int) { //接收任意数量的 int 作为参数
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...) //含有多个值的slice，作为参数使用（解包？）
}
