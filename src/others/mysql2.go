// mysql
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
	rows, err1 := db.Query("SELECT * FROM user_info ")
	if err1 != nil {
		panic(err1)
	}
	for rows.Next() {
		var id int
		var username string
		var department string
		var create_time string
		err = rows.Scan(&id, &username, &department, &create_time)
		fmt.Println(id, username, department, create_time)
	}

}
