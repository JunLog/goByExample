package main

import "fmt"

func main() {
	i := 1
	for i < 3 {//单条件循环
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j <= 9; j++ {//初始化；条件；后续
		fmt.Println(j)
	}
	for {//死循环
		fmt.Println("loop")
		break//跳出循环（return同效）
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue//跳过本轮循环
		}
		fmt.Println(n)
	}
}
