package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	//包前面加下横线代表只导入该包内的init方法
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	stmt1, err3 := db.Prepare("delete from user_info where id=?")
	if err3 != nil {
		panic(err3)
	}

	_, err4 := stmt1.Exec(4)
	if err4 != nil {
		panic(err4)
	}

	rows2, err6 := db.Query("SELECT * FROM user_info ")
	if err6 != nil {
		panic(err6)
	}
	for rows2.Next() {
		var id int
		var username string
		var department string
		var create_time string
		err = rows2.Scan(&id, &username, &department, &create_time)
		fmt.Println(id, username, department, create_time)
	}
}
