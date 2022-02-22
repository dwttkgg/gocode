package main

import (
	"fmt"

	"github.com/gomodule/redis"
)

//定义一个全局pool
var pool *redis.Pool

func init() {

	// 初始化pool连接池配置
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲数
		MaxActive:   0,   // 最大连接数  0表示不限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接代码，表示连接哪个redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func main() {
	// 从pool中取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "cat", "tom")
	if err != nil {
		fmt.Println(err)
	}
	r, err := redis.String(conn.Do("get", "cat"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)

	// pool.close()连接池关闭后，就无法正常使用
}
