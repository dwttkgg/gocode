package uitl

import (
	"fmt"
	"ginbook/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB() (err error) {
	addr := "root:123456@tcp(192.168.30.101:3310)/test"
	db, err = sqlx.Connect("mysql", addr)
	if err != nil {
		return err
	}
	// 最大连接
	db.SetMaxOpenConns(100)
	// 最大空闲
	db.SetMaxIdleConns(16)
	return
}
func List() (booklist []*model.Book, err error) {
	sqlstr := "select * from book"
	err = db.Select(&booklist, sqlstr)
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	return
}
func InsertBook(title string, price int64) (err error) {
	sqlStr := "insert into book(title,price) values (?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("插入失败", err)
		return
	}
	return
}
func DeleteBook(id int64) (err error) {
	sqlStr := "delete from  boook where id =?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除失败", err)
		return
	}
	return
}
