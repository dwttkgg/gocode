package main

import (
	"fmt"
	"gocode/test/gongchang/model"
)

//  使用工厂模式解决 model包中 struct 名称小写无法使用的问题
func main() {
	//stu1 := model.Student{"tang", 28, 22.3}
	stu1 := model.NewStudent("tang", 22, 22.2)
	fmt.Println(stu1)
	fmt.Println(stu1.Age, stu1.Score, stu1.GetName())
}
