package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mymap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	// sync    synchronized  mutex互斥
	lock sync.Mutex
)

func test(n int) {
	var res int = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	mymap[n] = res
	lock.Unlock()
}

// 使用goroutine来完成1-200的阶乘
// 使用 go build -race main.go 来查看是否有资源竞争
func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	//休眠10s 以免主线程结束导致协程无法执行
	time.Sleep(time.Second * 10)
	lock.Lock()
	for i, v := range mymap {
		fmt.Println(i, v)
	}
	lock.Unlock()
}
