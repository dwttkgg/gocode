package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("n", "ganyu", "请输入名称")
	age := flag.Int("age", 2000, "请输入年龄")
	flag.Parse()
	fmt.Println(*name, *age)
}
