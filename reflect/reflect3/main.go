package main

import (
	"fmt"
	"reflect"
)

type Pers struct {
	Name  string `json:"名称"`
	Age   int    `json:"年龄"`
	Skill string
	sc    float64
}

func (a Pers) Get(n1, n2 int) int {
	return n1 + n2
}

func (a Pers) Pri() {
	fmt.Println("开始打印")
	fmt.Println(a)
	fmt.Println("结束打印")
}
func (a Pers) Set(p Pers) {
	a.Name = p.Name
	a.Age = p.Age
	a.Skill = p.Skill
	a.sc = p.sc
}
func reflectApp(p interface{}) {
	rval := reflect.ValueOf(p)

	rkind := rval.Kind()
	rtype := reflect.TypeOf(p)
	if rkind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	// 获取结构体字段数量
	num := rtype.NumField()
	fmt.Println(num)
	for i := 0; i < num; i++ {
		// 获取对应字段的值
		fmt.Println(rval.Field(i))
		// 获取对应字段tag标签的值
		tag := rtype.Field(i).Tag.Get("json")
		if tag != "" {
			fmt.Println(tag)
		}
	}
	// 获取方法数  结构体的方法不能使用指针传递，否则查询不到方法
	numMeth := rval.NumMethod()
	fmt.Printf("struct has %d method \n", numMeth)

	rval.Method(1).Call(nil)
	// 调用结构体的第一个GET方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rval.Method(0).Call(params)
	fmt.Println(res[0].Int())
}
func main() {
	mon := Pers{
		Name:  "雷雨天",
		Age:   22,
		Skill: "ddd",
	}
	reflectApp(mon)
}
