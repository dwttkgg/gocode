package main

import "fmt"

func writeDate(intChan chan int) {
	for i := 1; i <= 500; i++ {
		intChan <- i
		fmt.Println("写入数据", i)
	}
	close(intChan)
}
func readDate(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读出数据", v)
	}
	exitChan <- true
	close(exitChan)
}
func main() {
	// 只有写进程，无读进程，make空间小于读入个数时会阻塞报错，
	// 读写进程都有，就算读写速度不一致也不会报错
	intChan := make(chan int, 500)
	exitChan := make(chan bool, 1)
	go writeDate(intChan)
	go readDate(intChan, exitChan)

	//循环读取 exitChan,来确保协程全部执行完成
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
