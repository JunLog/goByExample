package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 创建 临时目录和文件
// 程序结束后删除

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := ioutil.TempFile("", "sample")
	// 创建 并 打开 文件
	// 第一个参数传 "",将会在系统默认位置下创建
	check(err)

	fmt.Println("Temp file name:", f.Name())
	// 临时文件名，不完全等于创建时传进的第二个参数，是 参数作为前缀，后面部分自动生成
	// 这是为了确保并发调用时，生成不重复的文件名
	// 类 Unix 临时目录一般是： /tmp
	// Windows C:\Users\ADMINI~1\AppData\Local\Temp\

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4}) // 写入数据
	check(err)

	// 需要用到多个临时文件时，最好是创建一个文件夹
	dname, err := ioutil.TempDir("", "sampledir") 
	// 返回的是一个目录名，而不是打开的文件
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1") // 路径拼接
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)

}
