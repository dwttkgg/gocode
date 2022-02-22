package main

import "fmt"

type A interface {
	start()
	stop()
}

// go中实现了接口，需要实现所有接口中定义的方法，才是实现了该接口
type B struct{}

func (b B) start() {
	fmt.Println("B开始工作...")
}

func (b B) stop() {
	fmt.Println("B停止工作....")
}

type C struct{}

func (c C) start() {
	fmt.Println("C开始工作...")
}
func (c C) stop() {
	fmt.Println("C停止工作")
}

type D struct{}

func (d D) working(a A) {
	a.start()
	a.stop()
}
func main() {
	var d D = D{}
	var b B = B{}
	var c C = C{}
	d.working(b)
	d.working(c)
}
