package main

import (
	"fmt"
	"time"
)

func putChan(intChan chan int) {
	for i := 1; i <= 400000; i++ {
		intChan <- i
		//fmt.Println(i)
	}
	close(intChan)
}

func printNum(intChan chan int, printChan chan int, exitChan chan bool) {
	var flag bool
	//fmt.Println(flag)
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			printChan <- num
		}
	}
	fmt.Println("取完一个")
	exitChan <- true
}
func main() {
	intChan := make(chan int, 40000)
	printChan := make(chan int, 80000)
	// 退出协程
	exitChan := make(chan bool, 8)
	// 统计运行时间
	start := time.Now().Unix()
	//开启协程，并向intChan放入 1-8000的数据
	go putChan(intChan)
	//开启四个协程，从intChan取出数据，并判断是否为素数，是就放入printChan中
	for i := 0; i < 8; i++ {
		go printNum(intChan, printChan, exitChan)
	}
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(printChan)
		end := time.Now().Unix()
		fmt.Println("耗时:", end-start)
	}()

	for {
		_, ok := <-printChan
		if !ok {
			break
		}
		//fmt.Println("素数", res)
	}
	fmt.Println("main退出")
}
