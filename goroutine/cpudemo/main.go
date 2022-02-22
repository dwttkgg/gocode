package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 查看系统有多少个CPU
	num := runtime.NumCPU()
	fmt.Println(num)
	// 设置系统运行的cpu
	runtime.GOMAXPROCS(num - 1)

}
