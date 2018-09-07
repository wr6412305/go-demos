package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接mysql数据库
	db, err := sql.Open("mysql", "root:ljs199711@/world?charset=utf8")
	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("insert city set ID=?,Name=?,CountryCode=?,District=?,Population=?")
	checkErr(err)

	res, err := stmt.Exec(123456, "ljs", "ljs", "ljs", 123456)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
