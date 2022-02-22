package main

import "fmt"

var userId int
var passWd string

func main() {
	// 接受用户输入
	var key int
	// 判断是否继续显示菜单
	var show bool = true
	for show {
		fmt.Println("---------------海量用户通讯系统---------------")
		fmt.Println("\t\t\t1 登录系统")
		fmt.Println("\t\t\t2 注册用户")
		fmt.Println("\t\t\t3 退出系统")
		fmt.Println("请输入1-3")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录系统")
			show = false
		case 2:
			fmt.Println("注册用户")
			show = false
		case 3:
			fmt.Println("退出系统")
			show = false
		default:
			fmt.Println("输入不对，请重新输入1-3")
		}
	}

	if key == 1 {
		// 登录用户
		fmt.Println("请输入用户id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &passWd)
		err := login(userId, passWd)
		if err != nil {
			fmt.Println("登录失败", err)
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("用户注册逻辑")
	}
}
