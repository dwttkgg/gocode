package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err := sql.Open("mysql", "root:admin123@tcp(192.168.30.101:3310)/test")
	if err != nil {
		panic(err.Error())
	}
	Db.SetConnMaxIdleTime(100)
}
