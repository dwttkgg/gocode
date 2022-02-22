package main

import (
	"fmt"
)

func main() {
	var a int
	var b int

	for {
		fmt.Println("请输入年份")
		fmt.Scanln(&b)
		fmt.Println("请输入月份")
		fmt.Scanln(&a)
		switch a {
		case 1, 3, 5, 7, 8, 10, 12:
			fmt.Println("该月份有31天")

		case 4, 6, 9, 11:
			fmt.Println("该月份有30天")
		default:
			if b%4 == 0 && b%100 != 0 || b%400 == 0 {
				fmt.Println("该月份有29天")
			} else {
				fmt.Println("该月份有28天")
			}
		}
		continue
	}
}
