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
	//将id 改为自增  alter table  user_info id id(11) auto_increment
	stmt, err2 := db.Prepare("INSERT user_info  SET username=?,departname=?,create_time=?")
	if err2 != nil {
		panic(err2)
	}

	res, err3 := stmt.Exec("lisi", "technology", "2016-12-08")

	if err3 != nil {
		panic(err3)
	}
	id, err4 := res.LastInsertId() //插入此次成功标志

	if err4 != nil {
		panic(err4)
	}

	fmt.Println(id)
	//更新
	stmt, err5 := db.Prepare("update user_info set username=? where id = ?")
	if err5 != nil {
		panic(err5)
	}
	res, err7 := stmt.Exec("lisi", 22)
	if err7 != nil {
		panic(err7)
	}

	id, err8 := res.RowsAffected() //更新此次成功标志
	if err8 != nil {

		panic(err8)
	}
	fmt.Println(id)

}
