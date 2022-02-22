// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	// 99乘法表
// 	// var n int = 9
// 	// for i := 1; i <= n; i++ {
// 	// 	for j := 1; j <= i; j++ {
// 	// 		fmt.Printf("%d * %d = %d \t", i, j, i*j)
// 	// 	}
// 	// 	fmt.Println()
// 	// }

// 	// 随机数
// 	// 为了生成一个随机数，需要rand设置一个种子
// 	// time.Now().unix() 返回秒 time.Now().UnixNano() 返回纳秒
// 	// Intn左闭右开
// 	var count int = 0
// 	for {
// 		rand.Seed(time.Now().UnixNano())
// 		n := rand.Intn(100) + 1
// 		fmt.Println("n=", n)
// 		count++
// 		if n == 99 {
// 			break
// 		}
// 	}
// 	fmt.Println("生成99一共用了多少次", count)

// 	// break 可以配合标签使用，来退出到指定的循环
// 	// label1:    break label1
// 	// continue 和 break一致，作用跳出本次循环，也可使用标签

// 	// goto 配合标签指定跳转到某一行    return 普通函数跳出该函数，主函数，退出程序
// }
