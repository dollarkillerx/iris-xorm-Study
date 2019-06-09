/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-9
* Time: 上午10:02
* */
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, e := sql.Open("sqlite3", "./test.db")
	defer db.Close()
	if e != nil {
		panic(e.Error())
	}

	// insert
	stmt, e := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	if e != nil {
		panic(e.Error())
	}

	result, e := stmt.Exec("dollarkiller", "it", "2019-10-05")
	if e != nil {
		panic(e.Error())
	}
	i, _ := result.LastInsertId()
	fmt.Println("Index:",i)
}
