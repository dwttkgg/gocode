package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "d:/1.txt"
	// 一次性读取，无缓存，适用于不大的文件
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	// 将文件显示到屏幕
	//fmt.Printf("%v", content)
	fmt.Printf("%v", string(content))
}
