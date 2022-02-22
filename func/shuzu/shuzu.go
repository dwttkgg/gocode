package main

import (
	"fmt"
)

func main() {

	// var score [5]float64

	// for i := 0; i < len(score); i++ {
	// 	fmt.Println("请输入成绩")
	// 	fmt.Scanln(&score[i])
	// }
	// fmt.Println(score)
	// 四种初始化数组的方式
	var sz1 [3]int = [3]int{1, 2, 3}
	fmt.Println("sz1=", sz1)
	var sz2 = [3]int{1, 2, 3}
	fmt.Println("sz2=", sz2)
	var sz3 = [...]int{1, 2, 3}
	fmt.Println("sz1=", sz3)
	var sz4 = [...]int{1: 33, 0: 222, 3: 222}
	fmt.Println("sz1=", sz4)

	// 数组遍历 for range
	// index 下标，value 值，名称不固定，但当常用这个两个字词，如不需要可以用 _ 替换
	for index, value := range sz4 {
		fmt.Println(index, value)
	}
}
