package main

import "fmt"

func sayHel() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello", i)
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("gotest()", err)
		}
	}()
	var Maper map[int]string
	Maper[0] = "sss"
}
func main() {
	// 协程报错是导致整个程序崩溃
	// 使用defer + recover 来完善
	go sayHel()
	go test()

	for i := 0; i < 5; i++ {
		fmt.Println("main", i)
	}
}
