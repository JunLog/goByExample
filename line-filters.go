package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 行过滤器（line filter) 感觉就是处理命令行参数
// 读取 stdin 输入，处理，再输出到 stdout
// grep 和 sed 是常见的过滤器

func main() {
	// NewScanner 返回一个新的Scanner(序列?)，默认分割功能是ScanLines
	scanner := bufio.NewScanner(os.Stdin)

	// Scan 返回一个序列? 到达末尾或出现错误,返回false
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		// 程序没有进到这里
		fmt.Fprintln(os.Stderr, "error:", err) // Stderr 与 Stdin 和 Stdout 一样都会在命令行中打印出来
		// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
		// Fprintf 根据 format 参数生成格式化的字符串并写入 w[stream(流)指定的文件] 。返回写入的字节数和遇到的任何错误
		os.Exit(1)
	}
}

// echo hello>.\tmp\lines
// echo filter>>.\tmp\lines
// cmd 用 type ，linux 用 cat
// type .\tmp\lines | go run line-filters.go
