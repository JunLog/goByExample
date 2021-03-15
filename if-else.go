package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 是奇数")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	if num := 9; num < 0 {//条件语句前可以赋值
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
