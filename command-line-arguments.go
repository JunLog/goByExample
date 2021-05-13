package main

import (
	"fmt"
	"os"
)

//
func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg) // [./command-line-arguments a b c d]
	fmt.Println(argsWithoutProg) // [a b c d]
	fmt.Println(arg) // c
}

// 执行
// go build command-line-arguments.go
// command-line-arguments a b c d
