package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 读取客户端发送的信息
	defer conn.Close()

	// 循环读取客户端发消息
	for {
		buffer := make([]byte, 8096)
		n, err := conn.Read(buffer[:4])
		if err != nil || n != 4 {
			fmt.Println(err)
			return
		}
		fmt.Println(buffer)
	}
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("连接错误", err)
	}

	for {
		// 等待客户端连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		// 一旦连接成功，则启动一个协程处理连接请求
		go process(conn)
	}
}
