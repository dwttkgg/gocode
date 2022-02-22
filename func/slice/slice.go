package main

import (
	"fmt"
)

// slice 是引用类型，引用的数组
// 在内存中的布局主要是三部分，1 引用值的地址 2 slice的长度 3 slice的容量
// 类似一个一个struck
func main() {
	//切片使用的三种方式
	// 定义一个数组，用切片去引用创建好的数组  var inte [5]int = [5]int{17, 26, 31, 42, 33}  (数组事先存在，可见)
	// 通过make 去创建切片 make 切片名 []byte=make(byte[],len,cap)  type类型，len长度 cap容量 (make定义，由底层维护，不可见，只能通过slice访问)
	// 定义一个切片，直接指定具体数组 var slice []int = []int{17, 26, 31, 42, 33}

	var inte [5]int = [5]int{17, 26, 31, 42, 33}

	slice := inte[1:4]
	fmt.Println(slice)
	fmt.Println("", len(slice))
	fmt.Println("", cap(slice))
	var a []int 
	make a []int=make([]int,5,2)
	// 对slice进程扩容，使用append
	// 使用append增加具体元素
	slice = append(slice, 444, 333, 222)
	fmt.Println(slice)
	// 使用append增加切片
	slice = append(slice, slice...)
	fmt.Println(slice)
	// string 可以切片，但切片后不能修改string原始值，如果要修改，需要先转换成数组，修改后重新赋值给string
	// 涉及到中文赋值，需要使用 []rune 转换
	var str string = "hello world"
	var sli = str[6:]
	fmt.Println(sli)

	var byc []byte = []byte(str)
	byc[0] = 'A'
	str = string(byc)
	fmt.Println(str)

	//使用汉字
	var byc1 []rune = []rune(str)
	byc[0] = 'A'
	str = string(byc1)
	fmt.Println(str)

}
