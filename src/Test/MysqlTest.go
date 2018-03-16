package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/pholcus?charset=utf8 ")
	checkErr(err)
	//查询数据
	rows, err := db.Query("SELECT 用户ID FROM 新浪微博用户____用户详情")
	checkErr(err)
	for rows.Next() {
		//var name string
		var  用户ID string

		err = rows.Scan(&用户ID)
		checkErr(err)
		fmt.Println(用户ID)

	}

}
func checkErr(err error) {    if err != nil {        panic(err)    } }