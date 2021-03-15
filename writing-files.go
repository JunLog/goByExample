package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// cmd 删除一个文件
// del 文件名字

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("./tmp/dat1", d1, 0644)
	// 最后一个参数是:perm os.FileMode;是文件权限属性(Linux那一套)
	// 二进制表示,一个 9 位,
	// 第 1 位:文件属性,一般是 "-"(普通文件);还有是 "d" (目录)；
	// 第 2-4 位:文件所有者的权限 rwx (读/写/执行)
	// 第 5-7 位:文件所在用户组的权限 rwx (读/写/执行)
	// 第 8-10 位:其他人的权限 rwx (读/写/执行)
	// -rwxrwxrwx -> 0111111111(二进制) -> 0777(八进制，数字前加0)
	// 06444（8进制） -> -110 100 100
	check(err)

	f, err := os.Create("./tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)

	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() // 调用 Sync 将缓冲去的数据写入硬盘
	// Sync将文件的当前内容提交到稳定存储。通常，这意味着刷新文件系统中最近写入到磁盘的数据的内存副本。
	// 刷新缓冲 ？

	w := bufio.NewWriter(f) // 跟带缓冲的 Reader 一样，bufio也提供带缓冲的 Writer
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush() // 确保已将所有的缓冲操作应用于底层 writer
}

// cmd 查看文件命令 -> type 文件名 （类似 Linux 的 cat）
