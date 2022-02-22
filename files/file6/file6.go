package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	filename1 := "d:/1.txt"
	filename2 := "d:/2.txt"

	data, err := ioutil.ReadFile(filename1)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(filename2, data, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
