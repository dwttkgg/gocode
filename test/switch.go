// package utils

// import (
// 	"fmt"
// )

// // switch 和 if 比较
// // 1 如果判断的具体数值不多，而且符合整数，浮点数，字符，字符串这几种类型
// // 建议switch语句，简洁高效
// // 2 其他情况，对区间判断和结果位bool类型的判断，使用if，if的使用范围更广

// func utils() {
// 	var a byte
// 	fmt.Println("输入一个字母")
// 	fmt.Scanf("%c", &a)

// 	switch a {
// 	case 'a':
// 		fmt.Println("A")
// 	case 'b':
// 		fmt.Println("B")
// 	case 'c':
// 		fmt.Println("C")
// 	case 'd':
// 		fmt.Println("D")
// 	case 'e':
// 		fmt.Println("E")
// 	default:
// 		fmt.Println("other")
// 	}
// 	//fmt.Printf("a type %T  a=%q", a, a)

// 	var i int
// 	fmt.Println("输入成绩")
// 	fmt.Scanln(&i)
// 	switch {
// 	case i >= 60 && i <= 100:
// 		fmt.Println("合格")
// 	case i < 60:
// 		fmt.Println("不合格")
// 	default:
// 		fmt.Println("输入错误")
// 	}

// 	var m int
// 	fmt.Println("输入一个月份")
// 	fmt.Scanln(&m)
// 	switch m {
// 	case 3, 4, 5:
// 		fmt.Println("春季")
// 	case 8, 6, 7:
// 		fmt.Println("夏季")
// 	case 9, 10, 11:
// 		fmt.Println("秋季")
// 	case 12, 1, 2:
// 		fmt.Println("东季")
// 	}

// }
