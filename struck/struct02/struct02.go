package main

import (
	"fmt"
)

type Person struct {
	Name string
}

// 方法
// test与Person绑定，只能通过Person对象调用，不能通过其他类型和变量调用
// 所有的函数和方法，不管调用形式如何，最终决定是值拷贝还是地址拷贝的是定义方法时是跟上面类型绑定的
func (p Person) test() {
	fmt.Printf("他的名字是：%v \n", p.Name)
}

func (p Person) name() {
	fmt.Printf("%v是一个好人\n", p.Name)
}
func (p Person) getSum(n int, m int) int {
	return m + n
}
func main() {
	var p Person
	p.Name = "Tom"
	p.test()

	p.name()
	sum := p.getSum(10, 20)
	fmt.Println(sum)
}
