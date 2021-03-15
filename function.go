package main

import (
	"fmt"
)

func plus(a int, b int) int { //括号后面跟的是返回值类型
	return a + b
}

func plusPlus(a, b, c int) int { //连续参数同样类型，可以进声明最后一个参数
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3= ", res)
}
