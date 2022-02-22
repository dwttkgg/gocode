package model

import (
	"fmt"
	"strconv"
	"unsafe"
)

//  数据类型
//  复制当前行  shift + alt + 向下的光标
//  %d 字节  %T 数据类型 %c 字符 %v 按照变量的值输出
func main() {
	var i int = 10
	fmt.Println(i)

	//测试 int8 的范围  -128~127
	var j int8 = 12
	fmt.Println("j", j)

	// uint8 0-255
	var k uint8 = 255
	fmt.Println("k", k)

	// 查看数据的类型
	fmt.Printf(" i 的数据类型 %T \n", i)

	// 查看程序中变量占用字节的大小和数据类型
	// 在使用类型时，保小不保大原则，在不影响程序运行的情况下，尽量占用空间小的数据类型
	var n1 int8 = 12
	fmt.Printf("n1 的数据类型是 %T  n1 的字节大小是 %d", n1, unsafe.Sizeof(n1))

	// 单精度float32  双精度 float64
	// 浮点数据类型  符号位+指数位+尾数位
	// 浮点数可能会造成精度损失 64位比32位精度要大
	var f float64 = 5.1234
	f1 := 5.1234e2
	f2 := 5.1234e-2
	fmt.Println(f, f1, f2)
	// 字符型，go中的字符串是由字节组成，其他语言是由字符组成

	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1 =   c2 =", c1, c2)
	fmt.Printf("c1=%c c2=%c \n", c1, c2)

	//var c3 byte = '北'   ASCII 超出
	// 超过ASCII的字符，可以使用int来保存，如果需要输出码值，需要使用格式化输出printf
	var c3 int = '北'
	fmt.Printf("c3的ASCII的值=%d  c3的值%c\n", c3, c3)

	// 字符串数据类型  字符串一但赋值，不可更改，有两种表现形式，一种双引号，会识别转义字符，一种反引号 以字符串原生形式输出，包括换行和特殊字符，可以实现防止攻击，输出源代码等效果
	// 字符串数据拼接时，如果太长需要换行需要把 + 号放在上面

	// 基本数据类型的转换都要显示转换  var s int = float64(c)
	// 基本数据类型转换成string  使用 fmt.Sprintf('%参数',表达式)
	// 也可以使用func FormatBool(b bool) string ;func FormatInt(i int64, base int) string ;func FormatUint(i uint64, base int) string ;func FormatFloat(f float64, fmt byte, prec, bitSize int) string

	var in int64 = 123
	var bo bool = true
	var fl float64 = 123.22
	var in1 uint64 = 1231
	//var by byte='雷雨天'

	ini := strconv.FormatInt(in, 10)
	bob := strconv.FormatBool(bo)
	flf := strconv.FormatFloat(fl, 'f', 10, 64)
	inn := strconv.FormatUint(in1, 10)
	fmt.Printf("ini type %T ini=%q \n bob type %T bob=%q \n flf type %T flf=%q \n inn type %T inn=%q \n", ini, ini, bob, bob, flf, flf, inn, inn)

	var inc int = 111
	ina := strconv.Itoa(inc)
	fmt.Printf("inn type %T  inn=%q", ina, ina)

}
