package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s)) // 将字符串转成字节数组

	bs := h.Sum(nil) // 参数是用来追加额外的字节切片的（不需要）

	fmt.Println(s)
	fmt.Printf("%x\n", bs) // SHA1 常用16进制输出
}
