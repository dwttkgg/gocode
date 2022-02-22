package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}

// 自己创建处理器，实现ServeHttp方法
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的处理器")
}

func main() {
	myHandler := MyHandler{}
	server := http.Server{
		Addr:        "8080",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}
	//http.Handle("/myHandler", &myHandler)
	server.ListenAndServe()
}
