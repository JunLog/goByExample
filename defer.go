package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "error:%v\n", err)
		os.Exit(1)
	}
}

func main() {

	f := createFile("./tmp/defer.txt")
	// 貌似直接创建的话没有 tmp 文件夹 会出错
	defer closeFile(f) // 最后执行
	writeFile(f)
}
