package main

import (
	"fmt"
)

func main() {
	// 创建管道 并指明管道的类型
	// 如果需要在管道里面放入不同的数据类型可以使用  var intChan chan interface{}
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("类型是%T,默认值是%v \n", intChan, intChan)

	// 向管道写入数据 chan不能自动扩容，使用超过其容量将deadlock
	intChan <- 1
	num := 20
	intChan <- num
	intChan <- 222
	fmt.Printf("管道的值%v,管道的容量%v,管道的cap%v \n", intChan, len(intChan), cap(intChan))
	// 从管道取出数据 取完之后再取 将deadlock
	num1 := <-intChan
	fmt.Println(num1)
}
