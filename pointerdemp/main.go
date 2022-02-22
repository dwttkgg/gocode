package main

import (
	"fmt"
	"gocode/model"
)

// 值类型，都有对应的指针类型，*数据类型
// 值类型包括:基本数据类型 int,float,bool,string、数组和结构体struct
// 值类型，存储的变量和值一般放在内存的栈中，引用类型，值是地址，存放在堆中，地址指向的才行对应的值，放在栈中

// go env -w GO111MODULE=off          go env -w GOPATH=
func main() {
	var i int = 10
	fmt.Println("i 的地址", &i)
	//  创建要给指针指向对应值的地址，变量中存储的是一个地址，并且在内存中有自己的地址，使用*变量，可以获取到指向地址的值
	var po *int = &i
	fmt.Printf("po 的值 %v \npo 的地址 %v\n", po, &po)
	fmt.Printf("*po 的值 %v\n", *po)

	*po = 20
	fmt.Printf("i 的值改变%var \n", i)

	//变量名命名规则与其他一致
	//比如，你的变量名，常量名，函数名首字母大写，就可以被其他包访问，小写就是只能自己用，类似与Java的public和private

	fmt.Println(model.Str)

	// 取余  a % b = a - a / b * b

	var dd bool
	dd = 1 > 2 || 2 > 1
	fmt.Println("dd", dd)
}
