package main

import (
	"fmt"
	"gocode/exercises/work/customerManger/model"
	"gocode/exercises/work/customerManger/service"
)

type customerView struct {
	key             string // 接收用户输入
	loop            bool   // 是否循环输出
	customerService *service.CustomerService
}

// 添加客户
func (this customerView) add() {
	fmt.Println("---------------添加客户---------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮件：")
	email := ""
	fmt.Scanln(&email)

	//  创建一个新的customer，使用new2方法
	customer := model.NewCustomer2(name, gender, age, phone, email)

	if this.customerService.Add(customer) {
		fmt.Println("添加客户成功")
	} else {
		fmt.Println("添加失败")
	}

}

// 显示所有客户信息
func (this customerView) list() {
	customers := this.customerService.ListCustomer()
	fmt.Println("--------------客户列表-------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i, _ := range customers {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------客户列表完成------------\n\n")
}

// 显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("-------------客户管理软件-------------")
		fmt.Println("             1 添加客户")
		fmt.Println("             2 修改客户")
		fmt.Println("             3 删除客户")
		fmt.Println("             4 查询客户")
		fmt.Println("             5 退出")
		fmt.Println("请输入1-5")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			cv.list()
		case "5":
			cv.loop = false
		default:
			fmt.Println("你的输入有误")
		}

		if !cv.loop {
			break
		}
	}
}
func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCuntomerService()

	customerView.mainMenu()
}
