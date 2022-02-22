package main

import (
	"fmt"
	"reflect"
)

func ReflectTest1(b interface{}) {
	tval := reflect.ValueOf(b)
	rtype := reflect.TypeOf(b)

	rkind := tval.Kind()
	fmt.Println(rtype, tval, rkind)

	rb := tval.Interface()
	fmt.Println(rb)
}

func main() {
	var a = 1.2
	ReflectTest1(a)
}
