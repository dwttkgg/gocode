package main

import (
	"fmt"

	"../model"
)

func main() {
	user := &model.User{
		ID: 1,
	}
	u, _ := user.GetUserById()
	fmt.Println("查询结果", u)
}
