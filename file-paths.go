package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Linux : dir/file
// Windows : dir\file

func main() {
	p := filepath.Join("dir1", "dir2", "filename") // 构建一个层次结构的路径（跨操作系统）
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))       // 自动删除多余的分割符
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // 自动删除多余的目录

	fmt.Println("Dir(p):", filepath.Dir(p))   // 自动分割出路径的 目录
	fmt.Println("Base(p):", filepath.Base(p)) // 自动分割出路径的 文件
	dir, file := filepath.Split(p)            // 分割出 目录 和 文件
	fmt.Println("Splie(p):", dir, file)

	fmt.Println(filepath.IsAbs("dir/file")) // 判断是否是 绝对路径
	fmt.Println(filepath.IsAbs("C:/dir/file"))
	// 示例的 "/dir/file" 判断是绝对路径，Windows10 下判断不是

	filename := "config.json"

	ext := filepath.Ext(filename) // 分割出文件的 拓展名，文件后缀
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext)) // 文件名清除拓展名后的值

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	// Rel 返回 第二个路径 相对于 第一个路径 的 相对路径，相对路径不存在则返回错误
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	// Rel 返回 第二个路径 相对于 第一个路径 的 相对路径，相对路径不存在则返回错误
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
