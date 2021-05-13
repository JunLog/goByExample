package main

import (
	"fmt"
	"strconv"
)

// 数字解析
func main() {
	f, _ := strconv.ParseFloat("1.234", 64) //数字为64位
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	// 0 自动推断数字进制，64 表示整型数是以64位存储
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64) // 自动识别16进制
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64) // ParseUint 跟 ParseInt 类似，ParseUint 用于无符号数
	fmt.Println(u)

	k, _ := strconv.Atoi("135") // Atoi是10进制整型转化
	// Atoi等价于ParseInt(s, 10, 0)，转换为int类型
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
