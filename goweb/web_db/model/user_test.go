package model

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("User的测试用例，测试User的方法")
	//通过 t.Run来执行子测试函数
	t.Run("添加测试用户", testAddUser)
}

//如果函数名不是以Test开头，那么该函数默认不执行，我们可用将它设置成为一个子函数
func testAddUser(t *testing.T) {
	//utils.Db.Ping()
	fmt.Println("添加测试用户:")
	// user := &User{
	// 	Username: "dwt",
	// 	Password: "tkgg",
	// 	Email:    "annn",
	// }
	//user.AddUser(user.Username, user.Password, user.Email)
	user := &User{}
	//fmt.Println(user.Username, user.Password, user.Email)

	user.AddUser()
}

func testGetUserId(t *testing.T) {
	user := &User{
		ID: 1,
	}
	u, _ := user.GetUserById()
	fmt.Println("查询结果", u)
}
