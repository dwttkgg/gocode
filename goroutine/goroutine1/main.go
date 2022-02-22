package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 20; i++ {
		fmt.Println("test  hello world", i)
		time.Sleep(time.Second)
	}
}
func main() {
	go test() // 开启协程
	for i := 0; i < 10; i++ {
		fmt.Println("main hello world", i)
		time.Sleep(time.Second)
	}
}
