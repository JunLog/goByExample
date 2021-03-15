package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := os.Mkdir("subdir", 0755) // 当前目录下创建一个子目录
	// 0755 -> 0b111101101
	// ——111(所有者) 101(同组) 101(所有人) [rwx 可读、可写、可执行]
	check(err)

	//defer os.RemoveAll("subdir") // 删除整个目录（类似 rm -rf）

	// 一个创建 临时文件 的帮助函数
	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	// 创建一个有层级的目录，（类似 mkdir -p）
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := ioutil.ReadDir("subdir/parent")
	// 列出 目录内容，返回一个 os.fileInfo 类型的切片对象
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("subdir/parent/child") // 修改当前工作目录，(类似 cd)
	check(err)

	c, err = ioutil.ReadDir(".") // 列出 当前 目录内容
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..") // cd 回到最开始的地方
	check(err)

	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
	// 遍历一个目录及其所有字目录
	// 接收 一个路径 和 回调函数(用于处理访问到的每个目录和文件)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
