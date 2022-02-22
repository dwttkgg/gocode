package main

import (
	"fmt"
)

type Cat struct {
	name  string
	age   int
	color string
}

func main() {
	// 创建struct的四种方式
	// 方式1
	var cat Cat
	cat.name = "zz"
	cat.age = 18
	cat.color = "red"
	fmt.Println(cat)

	// 方式2
	cat2 := Cat{"cat2", 21, "d"}
	fmt.Println(cat2)

	// 方式3
	// cat3是一个指针 标准写法赋值 (*cat3).name = "cat3"
	// 但也可以写成 cat3.name="cat3" 因为go的设计者，为了程序员方便会在底层对 cat3.name = "cat3"进行处理
	// 会给p3加上取值运算 (*cat3).name = "cat3"
	var cat3 *Cat = new(Cat)
	(*cat3).name = "cat3"
	cat3.name = "cat333"
	(*cat3).age = 21
	(*cat3).color = "white"
	fmt.Println(cat3)

	// 方式4
	// var cat *Cat=&Cat{}
	//  原理同上
	var cat4 *Cat = &Cat{}
	(*cat4).name = "cat4"
	cat4.name = "cat444"
	fmt.Println(*cat4)
}
