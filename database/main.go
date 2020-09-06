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

type User struct {
	uid        int
	username   string
	departname string
	created    string
}

func main() {
	dbinfo := strings.Join([]string{dbusername, ":", dbpassword, "@tcp(", dbhostsip, ")/", dbname, "?charset=utf8"}, "")
	//连接数据库
	db, err := sql.Open("mysql", dbinfo)
	// fmt.Println(db)
	checkErr(err)

	// 插入，下面两种写法都可以
	//stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) values(?,?,?)")
	checkErr(err)
	result, err := stmt.Exec("李四", "销售", "2019-08-02")
	checkErr(err)

	//step4：处理sql操作后的结果
	lastInsertId, err := result.LastInsertId()
	rowsAffected, err := result.RowsAffected()
	fmt.Println("lastInsertId", lastInsertId)
	fmt.Println("影响的行数：", rowsAffected)

	// 修改
	result1, err := db.Exec("UPDATE userinfo SET username=?, departname=? WHERE uid=?", "王二狗", "行政部", 2)
	checkErr(err)

	lastInsertId1, err := result1.LastInsertId()
	rowsAffected1, err := result1.RowsAffected()
	fmt.Println("lastInsertId", lastInsertId1)
	fmt.Println("影响的行数：", rowsAffected1)

	/*
	   查询：处理查询后的结果：
	       思路一：将数据，存入到map中
	       思路二：创建结构体：
	*/
	rows, err := db.Query("SELECT uid,username,departname,created FROM userinfo")
	checkErr(err)

	//fmt.Println(rows.Columns())

	datas := make([]User, 0)
	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		if err := rows.Scan(&uid, &username, &departname, &created); err != nil {
			fmt.Println("获取失败")
		}

		user := User{uid, username, departname, created}
		datas = append(datas, user)
	}

	// 处理数据存放进map中
	//datas := make([] map[string] interface{}, 0)
	//for rows.Next() {
	//	var uid int
	//	var username string
	//	var departname string
	//	var  created string
	//	if err := rows.Scan(&uid, &username, &departname, &created); err != nil {
	//		fmt.Println("获取失败")
	//	}
	//	map1 := make(map[string] interface{})
	//	map1["uid"] = uid
	//	map1["username"] = username
	//	map1["departname"] = departname
	//	map1["created"] = created
	//
	//	datas = append(datas, map1)
	//}

	// 释放资源
	stmt.Close()
	db.Close()

	for _, v := range datas {
		fmt.Println(v)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
