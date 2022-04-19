package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// mysql连接
const (
	userName = "admin"
	pwd      = "9OjOmENBBSjG&t%q"
	ip       = "192.168.56.2"
	port     = "3306"
	database = "test_01"
)

func main() {
	// admin:pwd@tcp(192.168.56.2:3306)/table?charset=utf8
	db, err := sql.Open("mysql", userName+":"+pwd+"@tcp("+ip+":"+port+")/"+database+"?charset=utf8")
	if err != nil {
		fmt.Println("conn error = ", err)
		return
	}
	stmt, err := db.Prepare("select * from test_01.a_test_num where id = 1 ")
	if err != nil {
		fmt.Println("prepar err = ", err)
		return
	}
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("exec err = ", err)
		return
	}
	for rows.Next() {
		var id int
		var name sql.NullString
		var content sql.NullString
		var age int
		var ext sql.NullString
		var version int
		err := rows.Scan(&id, &name, &content, &age, &ext, &version)
		if err != nil {
			fmt.Println("scan err = ", err)
			return
		}
		fmt.Println("id = ", id, ", name = ", name.String, ", content = ", content.String, ", age = ", age, ", ext = ", ext.String, ", version = ", version)
	}
	defer func(db *sql.DB, stmt *sql.Stmt, rows *sql.Rows) {
		errdb := db.Close()
		if errdb != nil {
			fmt.Println("close db error = ", errdb)
		}
		errst := stmt.Close()
		if errst != nil {
			fmt.Println("close st = ", errst)
		}
		errRow := rows.Close()
		if errRow != nil {
			fmt.Println("close row = ", errRow)
		}
	}(db, stmt, rows)
}
