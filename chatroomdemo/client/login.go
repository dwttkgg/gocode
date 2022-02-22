package main

import (
	"encoding/json"
	"fmt"
	"gocode/chatroomdemo/common/message"
	"net"
)

// 接受账号密码，判断能否登录
func login(userId int, passWd string) (err error) {
	//fmt.Printf("用户名:=%v,密码:=%v", userId, passWd)
	// 1 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	// 2 通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3 创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = passWd

	// 4 将 loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("josn.Marsha", err)
		return
	}
	mes.Date = string(data)

	// 5 将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(err)
		return
	}

	return nil
}
