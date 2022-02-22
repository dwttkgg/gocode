package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}
type yaog struct {
	Name  string `json:"name"` // 增加tag  json:"name"将序列化后的json名称转换成小写
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

// 结构体之间可以转换，但是要几个要求
// 结构体字段的类型，名称，个数要完全一样
func main() {
	var a A
	var b B
	a = A(b)
	fmt.Println(a, b)
	yagao := yaog{"牛魔王", 500, "铁扇"}
	s, err := json.Marshal(yagao)
	if err != nil {
		fmt.Println("转换错误", err)
	}
	fmt.Println(string(s))
}
