package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	// migration()
	crud()
}

func migration(){
	db, err := gorm.Open("mysql", "root:LZWxm!@#123456@tcp(47.100.220.174:3306)/go_db?charset=utf8")
	checkErr(err)
	defer db.Close()

	// migration
	db.AutoMigrate(&Product{})
}

func crud(){

}

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}