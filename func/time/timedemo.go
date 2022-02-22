package main

import (
	"fmt"
	"time"
)

func main() {
	// 1 获取当前时间
	str1 := time.Now()
	fmt.Printf("%v,%T/n", str1, str1)

	// 2 获取年月日时分秒
	fmt.Println("年=%v", str1.Year())
	fmt.Println("月=%v", str1.Month())
	fmt.Println("月=%v", int(str1.Month()))
	fmt.Println("日=%v", str1.Day())
	fmt.Println("时=%v", str1.Hour())
	fmt.Println("分=%v", str1.Minute())
	fmt.Println("秒=%v", str1.Second())

	// 格式化时间第一种格式
	fmt.Printf("%d-%d-%d %d:%d:%d \n", str1.Year(), str1.Month(), str1.Day(), str1.Hour(), str1.Minute(), str1.Second())
	FMdate := fmt.Sprintf("%d-%d-%d %d:%d:%d \n", str1.Year(), str1.Month(), str1.Day(), str1.Hour(), str1.Minute(), str1.Second())
	fmt.Printf("%v", FMdate)

	// 格式化时间第二种格式 规定格式，必须写该时间
	fmt.Println(str1.Format("2006-01-02 15:04:05"))
	fmt.Println(str1.Format("15:04:05"))
	fmt.Println(str1.Format("2006/01/02"))

	// 时间常量
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )
	//  比如100小时 100 * time.Hour

	var i int = 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
		if i == 100 {
			break
		}
	}

	// unix时间戳 unixnano时间戳(一般作用获取随机数字)
	fmt.Printf("unix %v  \n  unixnano  %v", str1.Unix(), str1.UnixNano())

}
