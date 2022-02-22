package main

import (
	"fmt"
	"reflect"
)

func ReflectTest(b interface{}) {
	// 通过反射获取到传入数据的type ,kind值
	// 1 先获取 reflect.Type
	// rtype 为int  但是是reflect的int  不是普通类型的int
	// 如果要使用rtype进行计算需要使用 rtype.Int()转化成int才能继续计算，其他同理
	rtype := reflect.TypeOf(b)
	fmt.Println("rtype=", rtype)

	// 2 获取 reflect.Value
	rval := reflect.ValueOf(b)
	fmt.Printf("rval=%v,rval type is %T \n", rval, rval)

	// 将rtype转成interface{}
	iv := rval.Interface()
	// 将 interface{} 通过类型断言转换成需要使用的类型
	num2 := iv.(int)
	fmt.Printf("num2=%v,num2 type is %T \n", num2, num2)
}
func main() {
	// 对基本数据类型int的基本操作
	var num int = 100
	ReflectTest(num)
}
