package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// cmd 写入 文件
// 写入
// echo hello > dat
// 追加
// echo go >> dat

func main() {

	// 将文件内容读到内存中
	dat, err := ioutil.ReadFile("./tmp/dat")
	check(err)
	fmt.Print(string(dat))
	fmt.Println("=-==-=-=-==-=-=-=-")
	f, err := os.Open("./tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1) // 读五个字节
	check(err)
	fmt.Printf("前 %d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(8, 0)
	// 跳到一个已知的位置开始读
	// 第一个参数:偏移量
	// 第二个参数:从何处,0文件原点;1当前偏移量;2末尾
	// 返回一个新的偏移量

	// 第6位:(32)空格；第7位：(13)归位符CR；第8位：(10)换行符
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("跳后 %d bytes @ (当前起点：%d): ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(8, 0)
	// 第6位:(32)空格；第7位：(13)归位符CR；第8位：(10)换行符
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("跳后 %d bytes @ (当前起点：%d): %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0) // 没有往回跳的，但可以归位
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5) //bufio 缓冲 的 读取数据（提高小读操作）
	check(err)
	fmt.Printf("前 5 bytes: %s\n", string(b4))

	f.Close() // 这种操作应该在 Open 后立马使用 defer 来完成

}
