package main

import (
	"fmt"
)

func main() {
	i := 1
	fmt.Printf("值%v  类型%T  地址%v \n", i, i, &i)

	// new  n 是一个*int  值是一块空间的地址    改变开辟的地址的值使用指针即可  *n =11
	n := new(int)
	fmt.Printf("值%v  类型%T  地址%v 指向的值%v \n", n, n, &n, *n)
}
