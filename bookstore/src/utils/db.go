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
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/bookstore")//数据库设置在此更改
	if err != nil {
		panic(err.Error())
	}
}
