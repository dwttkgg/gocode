package main

import "fmt"

type Cat struct {
	Nmae string
	Age  int
}

func main() {
	var intChan chan interface{}
	intChan = make(chan interface{}, 10)
	cat1 := Cat{"tom", 111}
	cat2 := Cat{"tom23", 222}

	intChan <- cat1
	intChan <- cat2
	cat11 := <-intChan
	//cat21 := <-intChan
	//直接取出会报错，因为传入后为空接口类型
	//需要使用类型断言，将cat11 变成 Cat 类型，才能使用cat11.name
	a := cat11.(Cat)
	fmt.Println(a.Nmae)

}
