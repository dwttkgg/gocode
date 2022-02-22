package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("C:/Users/tkgga/Desktop/巡检.txt")

	if err != nil {
		fmt.Println("打开失败")
	}
	fmt.Println("文件内容", *file)

	defer file.Close()
	if err != nil {
		fmt.Println("文件关闭失败")
	}
	/*
		创建一个reader 默认带缓冲，大小为4096
		const{
			defaultBufSize=4096
		}
	*/
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)

	}
}
