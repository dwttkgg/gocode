package main

import "fmt"

// channel的关闭和遍历
func main() {
	intChan := make(chan int, 100)
	intChan <- 100
	intChan <- 200
	//管道关闭后，写入会报错
	close(intChan)
	//intChan <- 300
	// 可正常读入
	num1 := <-intChan
	fmt.Println("finnal", num1)
	// 遍历时如果没有关闭，则在读完所有值后，会deadlock
	for v := range intChan {
		fmt.Println(v)
	}
}
