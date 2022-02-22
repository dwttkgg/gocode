package main

import (
	"encoding/json"
	"fmt"
)

type ss struct {
	Name  string `json:"mont_name"`
	Age   int
	Score float64
}

func testStruct() {
	// 创建结构体
	peop := ss{
		Name:  "张三",
		Age:   22,
		Score: 22.2,
	}
	//	序列化
	data, err := json.Marshal(peop)
	if err != nil {
		fmt.Println("序列化失败", err)
	}
	// 输出序列化后的结果
	fmt.Println(string(data))
}
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{}, 10)
	a["name"] = "李四"
	a["age"] = 222
	a["scrote"] = 33.33
	dat, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败", err)
	}
	fmt.Println(string(dat))
}
func testSlice() {
	var sli []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 0)
	m1["name"] = "赵1"
	m1["age"] = 10
	m1["sc"] = 22.2
	sli = append(sli, m1)
	var m2 map[string]interface{}
	m2 = make(map[string]interface{}, 0)
	m2["name"] = "赵2"
	m2["age"] = 20
	m2["sc"] = 33.3
	sli = append(sli, m2)
	data, err := json.Marshal(sli)
	if err != nil {
		fmt.Println("序列化失败", err)
	}
	fmt.Println(string(data))
}
func main() {
	testStruct()
	testMap()
	testSlice()
}
