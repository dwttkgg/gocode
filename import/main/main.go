package main

import (
	"fmt"

	"../utils"
)

// go build 编译可执行文件，指定到main包
// go build -o bin/my.exe ... 可指定生产的文件位置和名称
func main() {
	utils.Sal()
	fmt.Println("this is main")
}
