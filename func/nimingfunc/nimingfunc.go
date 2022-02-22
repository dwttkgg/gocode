package main

import (
	"fmt"
	"strings"
)

// 定义全局匿名函数
var (
	func1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func addUpper() func(int) int {
	var n int = 10
	return func(i int) int {
		n += i
		return n
	}
}

// 闭包，判断后缀是不是.jpg 是则返回文件名，否则加上.jpg
func makesuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}
func main() {
	// 匿名函数，定义并使用
	a := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(a)
	// 定义匿名函数变量，多次使用
	f := func(n1 int, n2 int) int {
		return n1 - n2
	}

	fmt.Println(f(10, 20))
	fmt.Println(func1(10, 20))

	//闭包
	t := addUpper()
	fmt.Println(t(1))
	fmt.Println(t(2))
	fmt.Println(t(3))

	//闭包样例2
	bg := makesuffix(".jpg")
	fmt.Println(bg("new.txt"))
	fmt.Println(bg("new.jpg"))

}
