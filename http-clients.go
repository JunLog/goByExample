package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	// get 请求,创建了http.client 对象并调用请求
	// 使用 http.DefaultClient 对象及其默认设置
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 打印 HTTP response 状态
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	// 打印前五行
	for i := 0; scanner.Scan() && i < 5; i++ {
		// 按行读取
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
