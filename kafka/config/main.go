package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "../logs/beego.conf")
	if err != nil {
		fmt.Println("config.NewConfig is error:", err)
	}
	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server port is error:", err)
	}
	ip := conf.String("server::ip")

	fmt.Printf("the address is %v:%v\n", ip, port)
	lv := conf.String("logs::lever")
	fmt.Println("the logs leverl is :", lv)
	fmt.Println(lv)
}
