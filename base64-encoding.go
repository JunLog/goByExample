package main

import (
	b64 "encoding/base64" // 别名
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~" // 要解码的字符串

	sEnc := b64.StdEncoding.EncodeToString([]byte(data)) // 编码
	fmt.Println(sEnc)

	// 解码（base64 解码后缀为 +）
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// URL base64 编码
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	// URL base64 解码（URL base64 解码后缀为 -）
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
