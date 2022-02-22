package main

import (
	"fmt"
)

// 当执行到defer语句时，暂时不执行，会将defer后面的语句压入到栈中(为了方便称呼defer栈)
// 当函数执行完的时候，才会执行栈中的语句，按照先入后出的方式
// 在defer将语句放入栈中时，也会将对应的值放入栈中，后续修改栈中值不会改变
func sum(n1 int, n2 int) int {
	defer fmt.Println(n1)
	defer fmt.Println(n2)
	n1++
	fmt.Println(n1)
	fmt.Println(n1 + n2)
	return n1 + n2
}

func main() {
	sum(10, 20)
	fmt.Println("defer finished")
}
