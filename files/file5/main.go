package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个新文件并写入内容
	// 1 打开文件  d:/2.txt
	filename := "d:/2.txt"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭
	defer file.Close()
	// 准备写入 带缓存的writer
	str := "wo yao zhang gongzi A A A A A A \n"
	Writer := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		Writer.WriteString(str)
	}
	// 因为带缓存，所以需要刷到磁盘中 使用flush方法
	Writer.Flush()
}
