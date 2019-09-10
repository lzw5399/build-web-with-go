package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// insert()
	// update()
	// delete()
	query()
}

func insert() {
	db, err := sql.Open("mysql", "root:LZWxm!@#123456@tcp(47.100.220.174:3306)/go_db?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("test2", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func update() {
	db, err := sql.Open("mysql", "root:LZWxm!@#123456@tcp(47.100.220.174:3306)/go_db?charset=utf8")
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec("更新后的数据", 2)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func delete() {
	db, err := sql.Open("mysql", "root:LZWxm!@#123456@tcp(47.100.220.174:3306)/go_db?charset=utf8")
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func query() {
	db, err := sql.Open("mysql", "root:LZWxm!@#123456@tcp(47.100.220.174:3306)/go_db?charset=utf8")
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
