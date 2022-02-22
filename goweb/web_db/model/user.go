package model

import (
	"fmt"
	"gocode/goweb/web_db/utils"
)

// User 结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// 添加User
// func (user *User) AddUser() error {
// 	//写SQL语句
// 	sqlStr := "insert into users(username,PASSWORD,email) values(?,?,?)"

// 	// 预编译
// 	inStmt, err := utils.Db.Prepare(sqlStr)
// 	if err != nil {
// 		fmt.Println("预编译出错", err)
// 		return err
// 	}
// 	// 执行
// 	_, err = inStmt.Exec(user.Username, user.Password, user.Email)
// 	//_, err = inStmt.Exec("tkgg", "annn", "dwt")
// 	if err != nil {
// 		fmt.Println("执行出错", err)
// 		return err
// 	}
// 	return nil
// }
func (user *User) AddUser() error {
	//1.写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return err
	}
	//3.执行
	_, err2 := inStmt.Exec("admin", "123456", "admin@atguigu.com")
	if err2 != nil {
		fmt.Println("执行出现异常：", err2)
		return err
	}
	return nil
}

// 添加User 无预编译
func (user *User) AddUser2() error {
	//写SQL语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	// 执行
	_, err := utils.Db.Exec(sqlStr, user.Username, user.Password, user.Email)
	if err != nil {
		fmt.Println("执行出错", err)
		return err
	}
	return nil
}

// GetUserById
func (user *User) GetUserById() (*User, error) {
	sqlStr := "select id ,username,password,email from users where id =?"
	row := utils.Db.QueryRow(sqlStr, user.ID)
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		fmt.Println("查询报错", err)
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}
