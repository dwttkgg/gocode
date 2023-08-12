package main

import (
	"fmt"
	"net"
)

var Ip []string

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				Ip = append(Ip, ipnet.IP.String())

			}
		}
	}
	fmt.Println(Ip)
}
