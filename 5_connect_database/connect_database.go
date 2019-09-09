package main

import (
	"database/sql"
	"fmt"
	//"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:LZWxm!@#123456@tcp(47.100.220.174:3306)/go_db?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("test1", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
