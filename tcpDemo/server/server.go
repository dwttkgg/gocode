package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 关闭conn
	defer conn.Close()

	for {
		// 创建一个byte切片
		buf := make([]byte, 1024)
		// 1 等待客户端通过conn 发送消息
		// 2 如果客户端没有write  那么协程阻塞在这
		//fmt.Printf("服务器在等待%s客户端输入", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 3 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}

}
func main() {
	fmt.Println("服务器开始监听")
	// 表示使用的是tcp协议，监听8088端口
	listen, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	// 等待客户端连接
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端连接报错", err)

		} else {
			fmt.Println(conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
	//fmt.Printf("Listen success=%v", listen)
}
