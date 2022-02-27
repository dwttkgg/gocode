package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// go_path路径下执行go get github.com/garyburd/redigo/redis
func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	// 操作reids
	_, err = c.Do("set", "name", "china中国")
	if err != nil {
		fmt.Println("set err", err)
	}
	// redis自带类型转换，最好使用redis中的
	r, err := redis.String(c.Do("get", "name"))
	if err != nil {
		fmt.Println("set err", err)
	}
	fmt.Println(r)
}
