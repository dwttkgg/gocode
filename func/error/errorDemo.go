package main

import (
	"errors"
	"fmt"
)

func er() {
	// defer + recover 来捕获和处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n1 := 10
	n2 := 0
	res := n1 / n2
	fmt.Println(res)
}

// 自定义错误 使用errors.New 和 panic 内置函数
// errors.New("错误说明") 会返回一个error类型的值，表示一个错误
// panic内置函数 接受一个interface() 类型的值作为参数，可以接受error类型的变量并输出错误信息，退出程序
func readConf(strings string) (err error) {
	if strings == "config.ini" {
		return nil
	} else {
		return errors.New("文件名不匹配。。。")
	}
}
func test01(str string) {
	res := readConf(str)
	if res != nil {
		panic(res)
	}
}
func main() {
	er()
	test01("config.in")
	fmt.Println("main is running")
}
