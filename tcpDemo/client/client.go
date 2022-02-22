package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "30.15.160.46:8088")
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	//fmt.Println(conn)

	// 功能1  客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin)
	for {
		// 从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			return
		}
		// 发送给服务器
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn write err=", err)
		}
		fmt.Printf("发送了%v 字节的数据", n)

	}
}
