package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var (
	dbhostsip  = "10.211.55.13:3306"
	dbusername = "root"
	dbpassword = "renwoxing"
	dbname     = "test"
)

/*
事务操作是通过三个方法实现：
Begin()：开启事务
Commit()：提交事务（执行sql)
Rollback()：回滚
*/

func main() {
	dbinfo := strings.Join([]string{dbusername, ":", dbpassword, "@tcp(", dbhostsip, ")/", dbname, "?charset=utf8"}, "")
	//连接数据库
	db, err := sql.Open("mysql", dbinfo)
	// fmt.Println(db)
	checkErr1(err)
	//开启事务
	tx, _ := db.Begin()
	//提供一组sql操作
	var aff1, aff2 int64 = 0, 0
	result1, _ := tx.Exec("UPDATE account SET money=3000 WHERE id=?", 1)
	result2, _ := tx.Exec("UPDATE account SET money=2000 WHERE id=?", 2)
	//fmt.Println(result2)
	if result1 != nil {
		aff1, _ = result1.RowsAffected()
	}
	if result2 != nil {
		aff2, _ = result2.RowsAffected()
	}
	fmt.Println(aff1)
	fmt.Println(aff2)

	if aff1 == 1 && aff2 == 1 {
		//提交事务
		tx.Commit()
		fmt.Println("操作成功。。")
	} else {
		//回滚
		tx.Rollback()
		fmt.Println("操作失败。。。回滚。。")
	}

}

func checkErr1(err error) {
	if err != nil {
		panic(err)
	}
}
