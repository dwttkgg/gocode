package main

import "fmt"

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func fd(fun func(int, int) int, a1 int, a2 int) int {
	return fun(a1, a2)
}

//支持对返回值命名
func jisuanmyint(n1 int, n2 int) (sum int, min int) {
	sum = n1 + n2
	min = n1 - n2
	return
}

// 可变参数 args
// args是slice切片，通过args[index]来访问各个值
// 参数定义中,args一定是最后一位
func argss(n1 int, args ...int) (sum int) {
	sum = n1
	for i := 0; i < len(args); i++ {
		sum = sum + args[i]
	}
	return
}
func main() {
	fmt.Println(sum(10, 30))
	re1 := fd(sum, 50, 60)
	fmt.Println(re1)

	//自定义数据类型
	type myint int
	var num1 myint = 40
	// go认为 myint不是Int  所以将myint的值赋值给int时需要显式转换
	//var num2 int = int(num1)
	fmt.Println(num1)

	fmt.Println(jisuanmyint(40, 30))

	fmt.Println(argss(2, 3, 5, 3, -1))
}
