package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// 参数 ResponseWriter 是用于写入 HTTP 响应数据
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	// 将请求过来时的 req 内容返回
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v:%v\n", name, h)
		}
	}
}

func main() {
	// 将 handler 注册到服务器路由
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// 启动监听，nil 表示使用刚设的默认路由器
	http.ListenAndServe(":8090", nil)
}
