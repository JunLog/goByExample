package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
// 行过滤器（line filter)
// 读取 stdin 输入，处理，再输出到 stdout
// grep 和 sed 是常见的过滤器

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// echo hello>.\tmp\lines
// echo filter>>.\tmp\lines
// type .\tmp\lines | go run line-filters.go
