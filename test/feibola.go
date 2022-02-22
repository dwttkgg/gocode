package main

import (
	"fmt"
)

func feibona(n int) int {
	// var sum int = 1
	// var num int = 1
	// if n == 1 || n == 2 {
	// 	sum = 1
	// }
	// for i := 3; i <= n; i++ {
	// 	sum = sum + num
	// 	num = sum - num
	// }
	// return sum

	//递归方法
	// if n == 1 || n == 2 {
	// 	return 1
	// } else {
	// 	return feibona(n-1) + feibona(n-2)
	// }

	if n == 1 {
		return 3
	} else {
		return 2*feibona(n-1) + 1
	}

}

func main() {
	var n int
	fmt.Println("请输入一个数字")
	fmt.Scanln(&n)

	fmt.Println("斐波拉系数是", feibona(n))

}
