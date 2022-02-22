package main

func main() {
	// 管道可以声明为只读或者只写
	// 用处 在函数调用时，可以指定只读或者只写操作

	// 1 在默认情况下 管道是双向的
	//var chan1 chan int  //可读可写的双向管道
	// 2 声明为只写
	var chan1 chan<- int
	chan1 = make(chan int, 0)
	chan1 <- 1
	//fmt.Println(<-chan1)  // 报错，不能读取

	// 3 声明为只写
	//var chan2 <-chan int

}
