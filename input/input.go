package main

import (
	"fmt"
)

func main() {
	// var name string
	// var age int
	// var sal float64
	// var isPass bool
	// fmt.Println("请输入姓名")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)
	// fmt.Println("请输入工资")
	// fmt.Scanln(&sal)
	// fmt.Println("是否通过")
	// fmt.Scanln(&isPass)

	// fmt.Println(name, age, sal, isPass)

	// 大于等于0的原码，补码，反码不变    运算都是通过补码计算,正数可直接计算，负数补码计算后需要转换成反码在转换成原码得到最终结果
	// 负数的 -1 源码  10000001 反码 11111110 补码 11111111 第一位表示符号0正1负，反码除了符号位不变其他取反，补码是反码加1
	// 按位与&  两位全为1 结果位1 否则是0
	// 按位或|  两位有一个位1 结果位1 否则为0
	// 按位异或^ 两位一个位0一个位1 结果位1 否则位0
	// 测试  2 0000 0010   3 0000 0011  2&3 0000 0010 2|3 0000 0011 2^3 0000 0001
	fmt.Println(2 & 3)
	fmt.Println(2 | 3)
	fmt.Println(2 ^ 3)
	// -2 1000 0010 反码 1111 1101 补码 1111 1110
	fmt.Println(-2 ^ 2) // 补码 1111 1100 -> 反码 1111 1011 -> 原码 1000 0100

	// >>  << 移位运算符
	// 右移运算符 >> 低位溢出 符号位不变 符号位补溢出高位
	// 左移运算符 << 符号位不变，低位补0
	fmt.Println(1 >> 2) // 0000 0001 -> 0000 0000
	fmt.Println(1 << 2) // 0000 0001 -> 0000 0100

	// switch 分支 case 可以匹配多个表达式 且不用加breake  switch 表达式 {    case 表达式1,表达式2: ...   case 表达式3，表达式4：... }
	// switch 的数据类型和 case的数据类型一致   switch 也是做范围判断 switch { case scoge>90: ... case scoge >60 && scoge < 90 :...}
	// switch 穿透 在case 语句块中加上 fallthrough ,将会继续执行下一个case当中的内容 默认只能穿透一层

	var n1 int
	fmt.Println("输入n1")
	fmt.Scanln(n1)
	switch n1 {
	case 20, 10, 5:
		fmt.Println("this  is  ok ")
	case 2, 6, 3:
		fmt.Println("sb")
	default:
		fmt.Println("没有匹配到")
	}
}
