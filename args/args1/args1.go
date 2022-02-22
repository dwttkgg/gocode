package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数有几个", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("第%v个参数是%v\n", i, v)
	}
}
