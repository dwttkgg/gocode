package main

import (
	"fmt"
	"time"
)

func main() {
	// 使用select 可以解决从管道取数据的阻塞问题

	// 定义一个管道 int 10
	chanint := make(chan int, 10)
	for i := 0; i < 10; i++ {
		chanint <- i
	}
	// 定义一个管道 string 5
	chanString := make(chan string, 5)
	for i := 0; i < 5; i++ {
		chanString <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时，如果不关闭会阻塞造成死锁
	// 问题 在实际开发中 有时会不确定什么时候关闭
	// 可以使用select 解决
	for {

		select {
		case v := <-chanint:
			fmt.Printf("从chanint中取得数据%d\n", v)
		case v := <-chanString:
			fmt.Printf("从stringint中取得数据%s\n", v)
		default:
			fmt.Println("无法取到数据，请重新添加逻辑")
			time.Sleep(time.Second)
		}
	}
}
