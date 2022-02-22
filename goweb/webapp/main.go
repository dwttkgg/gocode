package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!", r.URL.Path)
}
func main() {
	// 使用自己创建的多路复用器来处理请求
	mux := http.NewServeMux()
	//http.HandleFunc("/", handler)
	mux.HandleFunc("/", handler)

	// 创建路由
	//http.ListenAndServe(":8088", nil)
	http.ListenAndServe(":8088", mux)
}
