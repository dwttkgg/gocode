package main

import "fmt"

type Student struct{}

// 编写一个类型判断函数
func typePanduan(items ...interface{}) {
	for i, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个数据类型是bool,值是%v \n", i, v)
		case float32, float64:
			fmt.Printf("第%v个数据类型是float,值是%v \n", i, v)
		case int, int64, int32:
			fmt.Printf("第%v个数据类型是int,值是%v \n", i, v)
		case string:
			fmt.Printf("第%v个数据类型是string,值是%v \n", i, v)

		case Student:
			fmt.Printf("第%v个数据类型是Student,值是%v \n", i, v)
		case *Student:
			fmt.Printf("第%v个数据类型是*Student,值是%v \n", i, v)
		default:
			fmt.Printf("第%v个数据类型不确定,值是%v \n", i, v)
		}
	}
}
func main() {
	var n1 float32 = 1.1
	var n2 int64 = 6
	name := "tim"
	n3 := true

	stu := Student{}
	stu2 := &Student{}

	typePanduan(n1, n2, n3, name, stu, stu2)
}
