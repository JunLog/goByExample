package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s)) // 将字符串转成字节数组

	bs := h.Sum(nil) 
	// Sum() 得到最终的散列值的字符切片
	// Sum()的参数是用来追加额外的字节切片的（一般不需要）

	fmt.Println(s)
	fmt.Println(bs)
	fmt.Printf("%x\n", bs) // SHA1 常用16进制输出
}
